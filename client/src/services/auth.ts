import type { AuthSchema } from "../schemas/auth";
import { ErrorHandler } from "../lib/errors";

/**
 * @class Handler of all authentication requests for the API.
 */
export class AuthService {
    private static url: string = import.meta.env.PUBLIC_API_AUTH;


    /**
     * Request API for authentication.
     * 
     * @param {AuthSchema} body Data required for authentication.
     */
    static async login(body: AuthSchema): Promise<{ status: (number | null), info: (any | ErrorHandler) }> {
        try {
            const response = await fetch(`${this.url}/auth/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(body)
            });
            const token = await response.json();
            if (response.status === 200) return { status: response.status, info: token }
            else return { status: response.status, info: new ErrorHandler(response.status, response.statusText) }
        } catch (err) {
            return { status: null, info: new ErrorHandler(500, (err as Error).name + ": " + (err as Error).message) }
        }
    }
}