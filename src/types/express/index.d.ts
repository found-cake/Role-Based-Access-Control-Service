import { JwtPayload } from 'jsonwebtoken';
import { ApiResponse } from '../api/response';

declare global {
    namespace Express {
        interface Request {
            user?: string | JwtPayload;
        }
        interface Response {
            json<T = any>(body: ApiResponse<T>): this;
        }
    }
}
