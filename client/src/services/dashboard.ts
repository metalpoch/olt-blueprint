import { ErrorHandler } from "../lib/errors";


export class DashboardService {
    private static url: string = import.meta.env.PUBLIC_API_CORE;


    static async gettrafficbystate(month : string): Promise<{ status: (number | null), info: (any | ErrorHandler) }> {
        try {
            const response = await fetch(`${this.url}/traffic/state/${month}`,{
                method: 'GET'
            });
            console.log(response);
            if (response.ok) return { status: response.status, info: await response.json() };
            else return { status: response.status, info: new ErrorHandler(response.status, response.statusText) };
        } catch(err) {
            return { status: null, info: new ErrorHandler(500, (err as Error).name + ": " + (err as Error).message) };
        }

    }

    static async gettrafficbystatetopN(month : string, n : number): Promise<{ status: (number | null), info: (any | ErrorHandler) }> {
        try {
            const response = await fetch(`${this.url}/traffic/state/${month}/${n}`,{
                method: 'GET'
            });
            if (response.ok) return { status: response.status, info: await response.json() };
            else return { status: response.status, info: new ErrorHandler(response.status, response.statusText) };
        } catch(err) {
            return { status: null, info: new ErrorHandler(500, (err as Error).name + ": " + (err as Error).message) };
        }

    }


    static async gettrafficbyodn(month : string): Promise<{ status: (number | null), info: (any | ErrorHandler) }> {
        try {
            const response = await fetch(`${this.url}/traffic/odn_d/${month}`,{
                method: 'GET'
            });
            if (response.ok) return { status: response.status, info: await response.json() };
            else return { status: response.status, info: new ErrorHandler(response.status, response.statusText) };
        } catch(err) {
            return { status: null, info: new ErrorHandler(500, (err as Error).name + ": " + (err as Error).message) };
        }

    }

}