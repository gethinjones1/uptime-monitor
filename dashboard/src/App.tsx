import React, { useState, useEffect, createContext, useContext } from 'react';
import {
  BrowserRouter,
  Routes,
  Route,
  Navigate,
} from 'react-router-dom';
import { CLILogin } from './components/CLILogin';

import { Dashboard } from './components/Dashboard';
import { LoginScreen } from './components/LoginScreen';
import { SignUpScreen } from './components/SignupScreen';

// Auth context to hold login state
type AuthContextType = {
  token: string | null;
  login: (token: string) => void;
  logout: () => void;
};
const AuthContext = createContext<AuthContextType | null>(null);

export const useAuth = () => {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error('useAuth must be used within AuthProvider');
  return ctx;
};

const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem('authToken'));

  const login = (newToken: string) => {
    setToken(newToken);
    localStorage.setItem('authToken', newToken);
  };
  const logout = () => {
    setToken(null);
    localStorage.removeItem('authToken');
  };

  return (
    <AuthContext.Provider value={{ token, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// Protected route wrapper
const PrivateRoute: React.FC<{ children: JSX.Element }> = ({ children }) => {
  const { token } = useAuth();
  return token ? children : <Navigate to="/login" replace />;
};

// Public route wrapper: redirect authenticated users
const PublicRoute: React.FC<{ children: JSX.Element }> = ({ children }) => {
  const { token } = useAuth();
  return token ? <Navigate to="/dashboard" replace /> : children;
};

const App: React.FC = () => {
  const auth = useAuth();

  // Handlers call API and then login
  const handleLogin = async ({ username, password }: { username: string; password: string }) => {
    const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password }),
    });
    if (!res.ok) throw new Error('Login failed');
    const data = await res.json();
    auth.login(data.token);
  };

  const handleSignUp = async ({ firstName, lastName, email, password }: any) => {
    const res = await fetch('http://localhost:8080/signup', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ firstName, lastName, email, password }),
    });
    if (!res.ok) throw new Error('Sign up failed');
    const data = await res.json();
    auth.login(data.token);
  };

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/cli-login" element={<CLILogin />} />

        <Route
          path="/login"
          element={

            <LoginScreen onLogin={handleLogin} />
          }
        />
        <Route
          path="/signup"
          element={
            <PublicRoute>
              <SignUpScreen onSignUp={handleSignUp} />
            </PublicRoute>
          }
        />
        <Route
          path="/dashboard"
          element={
            <PrivateRoute>
              <Dashboard />
            </PrivateRoute>
          }
        />
        {/* root and any other unknown paths */}
        <Route
          path="/"
          element={
            auth.token ? (
              <Navigate to="/dashboard" replace />
            ) : (
              <Navigate to="/login" replace />
            )
          }
        />
        <Route
          path="*"
          element={
            auth.token ? (
              <Navigate to="/dashboard" replace />
            ) : (
              <Navigate to="/login" replace />
            )
          }
        />
      </Routes>
    </BrowserRouter>
  );
};

export default () => (
  <AuthProvider>
    <App />
  </AuthProvider>
);
