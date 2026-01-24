import fetch from 'node-fetch';

const BASE_URL = 'http://localhost:3000/auth';

const testAuth = async () => {
    // Register
    console.log('Testing Register...');
    const registerResponse = await fetch(`${BASE_URL}/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            email: 'test@example.com',
            password: 'password123',
            name: 'Test User'
        })
    });
    const registerData = await registerResponse.json();
    console.log('Register Response:', registerData);

    if (!registerData.success && registerData.error?.details?.includes('Unique constraint failed')) {
        console.log('User already exists, proceeding to login...');
    }

    // Login
    console.log('\nTesting Login...');
    const loginResponse = await fetch(`${BASE_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            email: 'test@example.com',
            password: 'password123'
        })
    });
    const loginData = await loginResponse.json();
    console.log('Login Response:', loginData);
};

testAuth().catch(console.error);
