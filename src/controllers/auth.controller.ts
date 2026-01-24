import { Request, Response } from 'express';
import { AuthService } from '../services/auth.service';

const authService = new AuthService();

export const register = async (req: Request, res: Response) => {
    try {
        const { email, password, name } = req.body;
        const { user, token } = await authService.register({ email, password, name });
        res.status(201).json({
            success: true,
            data: { user, token },
            message: 'User registered successfully',
        });
    } catch (error: any) {
        res.status(400).json({
            success: false,
            error: { code: 400, details: error.message },
        });
    }
};

export const login = async (req: Request, res: Response) => {
    try {
        const { email, password } = req.body;
        const result = await authService.login({ email, password });

        if (!result) {
            return res.status(401).json({
                success: false,
                error: { code: 401, details: 'Invalid credentials' },
            });
        }

        res.status(200).json({
            success: true,
            data: result,
            message: 'Login successful',
        });
    } catch (error: any) {
        res.status(500).json({
            success: false,
            error: { code: 500, details: error.message },
        });
    }
};
