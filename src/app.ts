import express, { Express } from 'express';
import dotenv from 'dotenv';
import { PrismaClient } from '@prisma/client';
import { Pool } from 'pg';
import { PrismaPg } from '@prisma/adapter-pg';

dotenv.config();

const connectionString = process.env.DATABASE_URL;
const pool = new Pool({ connectionString });
const adapter = new PrismaPg(pool);
export const prisma = new PrismaClient({ adapter });

const app: Express = express();

import authRoutes from './routes/auth.routes';

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(express.static('public'));

app.use('/auth', authRoutes);

app.get('/health', (req, res) => {
    res.status(200).json({ status: 'ok' });
});

export default app;
