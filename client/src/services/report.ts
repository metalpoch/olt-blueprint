import { ErrorHandler } from "../lib/errors";

/**
 * @class Handler of all authentication requests for the API.
 */
export class ReportService {
    private static url: string = process.env.PUBLIC_API_REPORT ?? import.meta.env.PUBLIC_API_REPORT;

    /**
     * Request API for the upload of a file.
     * 
     * @param {file} file File to be uploaded.
     * @param {id} id ID of the file.
     * @param {category} category Category of the file.
     */
    static async uploadFile(file: any, id: number, category: string = "Reporte Regional"): Promise<{ status: (number | null), info: (any | ErrorHandler) }> {
        try {
            let data = new FormData();
            data.append("category", category);
            data.append("user_id", id.toString());
            data.append("file", file);

            const response = await fetch(`${this.url}/report`, {
                method: "POST",
                body: data
            });
            if ((response.status === 201) || (response.status === 200)) return { status: response.status, info: null }
            else return { status: response.status, info: new ErrorHandler(response.status, response.statusText) }
        } catch (err) {
            return { status: null, info: new ErrorHandler(500, (err as Error).name + ": " + (err as Error).message) }
        }
    }
}
