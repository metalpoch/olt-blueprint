import json
from os import getenv

from dotenv import load_dotenv
from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware

from src import model
from src import constants
from src.database import Postgres
from src.libs.chatbox import AI
from src.libs.linear_regression import RegressionLineal
from src.libs.osm import Openstreetmap
from src.libs.tracking import Telegram
from src.utils import change, execute


load_dotenv()


telegram = Telegram(getenv("TELEGRAM_BOT_ID", ""), getenv("TELEGRAM_CHAT_ID", ""))
app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Permite todos los orígenes
    allow_credentials=True,
    allow_methods=["*"],  # Permite todos los métodos HTTP
    allow_headers=["*"],  # Permite todos los headers
)
db = Postgres(getenv("URI", ""))

SQL_SYSTEM_PROMPT = f"{constants.PROMPT_1}\n```csv\n{db.csv_schemas()}\n```\n{constants.PROMPT_2}\n{constants.PROMPT_3}"


@app.post("/trend")
async def linear_regression(sysname: str, future_month: int):
    id = execute.sys_name_to_id(db, sysname)
    trends = []
    response = execute.trends(db, id)
    for trend in response:
        new_trend = change.create_trend(
            trend[0], trend[1], trend[4], trend[3], trend[2]
        )
        trends.append(new_trend)
    out, in_ = RegressionLineal.run_procress(trends, future_month)
    response_dict = []
    for i in trends:
        response_dict.append({"month": i.date.month, "in": i.in_, "out": i.out})
    return {"months": response_dict, "out_trend": out, "in_trend": in_}


@app.post("/chatbox")
async def chatbox(request: model.QueryAI) -> dict:
    ai_sql = AI(model="gemma2:2b", system_prompt=SQL_SYSTEM_PROMPT)
    ai_output = AI(model="gemma2:2b", system_prompt=constants.REPHRASE_PROMPT)

    try:
        msg = ai_sql.query(request.message)
        sql = ai_sql.sql_extract(msg)

        res, err = db.execute_sql(sql)
        if err is not None:
            raise HTTPException(status_code=400, detail=msg)

        output = ai_output.query(f"{res}")
        return {"output": output}

    except BaseException as e:
        raise HTTPException(status_code=400, detail=str(e))


@app.post("/telegram")
async def tracking(req: model.Telegram) -> dict:
    try:
        res = telegram.send_message(
            module=req.module, event=req.event, category=req.category, msg=req.message
        )

    except BaseException as e:
        raise HTTPException(status_code=400, detail=str(e))

    else:
        return {"response": res}


@app.get("/location")
async def location(latitude: float, longitude: float) -> dict:
    try:
        res = Openstreetmap(latitude, longitude).location()

    except BaseException as e:
        raise HTTPException(status_code=400, detail=str(e))

    else:
        return res
