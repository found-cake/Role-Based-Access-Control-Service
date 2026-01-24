import bcrypt from 'bcrypt';
import { prisma } from '../app';
import { generateToken } from '../utils/jwt';
import { User, Prisma } from '@prisma/client';

export class AuthService {
    async register(data: Prisma.UserCreateInput): Promise<{ user: User; token: string }> {
        const hashedPassword = await bcrypt.hash(data.password, 10);
        const user = await prisma.user.create({
            data: {
                ...data,
                password: hashedPassword,
            },
        });

        const token = generateToken({ id: user.id, email: user.email, role: user.role });
        return { user, token };
    }

    async login(data: Pick<Prisma.UserCreateInput, 'email' | 'password'>): Promise<{ user: User; token: string } | null> {
        const user = await prisma.user.findUnique({ where: { email: data.email } });
        if (!user) return null;

        const isPasswordValid = await bcrypt.compare(data.password, user.password);
        if (!isPasswordValid) return null;

        const token = generateToken({ id: user.id, email: user.email, role: user.role });
        return { user, token };
    }
}
