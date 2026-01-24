const API_URL = 'http://localhost:3000/auth';

// Helper to get token
const getToken = () => localStorage.getItem('token');

// Helper to set user session
const setSession = (token, user) => {
    localStorage.setItem('token', token);
    localStorage.setItem('user', JSON.stringify(user));
};

// Helper to clear session
const clearSession = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    updateNav();
};

// Update Navbar based on auth state
const updateNav = () => {
    const token = getToken();
    const navLogin = document.getElementById('nav-login');
    const navRegister = document.getElementById('nav-register');
    const navLogout = document.getElementById('nav-logout');
    const navDashboard = document.getElementById('nav-dashboard');

    if (token) {
        if (navLogin) navLogin.classList.add('hidden');
        if (navRegister) navRegister.classList.add('hidden');
        if (navLogout) navLogout.classList.remove('hidden');
        if (navDashboard) navDashboard.classList.remove('hidden');
    } else {
        if (navLogin) navLogin.classList.remove('hidden');
        if (navRegister) navRegister.classList.remove('hidden');
        if (navLogout) navLogout.classList.add('hidden');
        if (navDashboard) navDashboard.classList.add('hidden');
    }
};

// Handle Logout
const handleLogout = (e) => {
    e.preventDefault();
    clearSession();
    window.location.href = 'index.html';
};

// Handle Login Form
const handleLogin = async (e) => {
    e.preventDefault();
    const form = e.target;
    const formData = new FormData(form);
    const data = Object.fromEntries(formData.entries());
    const errorMsg = document.getElementById('error-message');

    try {
        const response = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        });
        const result = await response.json();

        if (result.success) {
            setSession(result.data.token, result.data.user);
            window.location.href = 'dashboard.html';
        } else {
            errorMsg.textContent = result.error.details || 'Login failed';
        }
    } catch (error) {
        errorMsg.textContent = 'An error occurred. Please try again.';
        console.error(error);
    }
};

// Handle Register Form
const handleRegister = async (e) => {
    e.preventDefault();
    const form = e.target;
    const formData = new FormData(form);
    const data = Object.fromEntries(formData.entries());
    const errorMsg = document.getElementById('error-message');

    try {
        const response = await fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        });
        const result = await response.json();

        if (result.success) {
            setSession(result.data.token, result.data.user);
            window.location.href = 'dashboard.html';
        } else {
            errorMsg.textContent = result.error.details || 'Registration failed';
        }
    } catch (error) {
        errorMsg.textContent = 'An error occurred. Please try again.';
        console.error(error);
    }
};

// Populate Dashboard
const populateDashboard = () => {
    const userStr = localStorage.getItem('user');
    if (!userStr) return;

    const user = JSON.parse(userStr);
    const nameEl = document.getElementById('profile-name');
    const emailEl = document.getElementById('profile-email');
    const roleEl = document.getElementById('profile-role');
    const greetingEl = document.getElementById('user-greeting');

    if (nameEl) nameEl.textContent = user.name;
    if (emailEl) emailEl.textContent = user.email;
    if (roleEl) roleEl.textContent = user.role.toUpperCase();
    if (greetingEl) greetingEl.textContent = `Hi, ${user.name}`;
};

// Init
document.addEventListener('DOMContentLoaded', () => {
    updateNav();

    const logoutBtn = document.getElementById('nav-logout');
    if (logoutBtn) logoutBtn.addEventListener('click', handleLogout);

    const loginForm = document.getElementById('login-form');
    if (loginForm) loginForm.addEventListener('submit', handleLogin);

    const registerForm = document.getElementById('register-form');
    if (registerForm) registerForm.addEventListener('submit', handleRegister);

    if (window.location.pathname.includes('dashboard.html')) {
        populateDashboard();
    }
});
