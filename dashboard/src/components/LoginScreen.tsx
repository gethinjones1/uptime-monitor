import React, { useState } from "react";
import { useSearchParams, useNavigate } from "react-router-dom";

export const LoginScreen: React.FC<{ onLogin: (credentials: { username: string; password: string }) => void }> = ({ onLogin }) => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    const [searchParams] = useSearchParams();
    const navigate = useNavigate();
    const redirect = searchParams.get("redirect") || "/dashboard";

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError(null);
        setLoading(true);
        try {
            await onLogin({ email, password });
            navigate(redirect);   // 🚀 Redirect after login
        } catch (err: any) {
            setError(err.message || "Login failed");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="flex min-h-screen items-center justify-center bg-white font-[Manrope,_Noto_Sans,_sans-serif]">
            <div className="w-full max-w-md space-y-8 p-10 rounded-2xl shadow-lg">
                <div className="text-center">
                    <h1 className="text-3xl font-bold text-[#111418]">SiteGuard Login</h1>
                    <p className="mt-2 text-sm text-[#637588]">Access your dashboard</p>
                </div>
                {error && <div className="text-red-500 text-sm text-center">{error}</div>}
                <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
                    <div className="space-y-4">
                        <div>
                            <input
                                id="username"
                                type="text"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                required
                                placeholder="Username or email"
                                className="w-full px-4 py-3 rounded-xl bg-[#f0f2f4] text-base text-[#111418] placeholder-[#637588] focus:outline-none"
                            />
                        </div>
                        <div>
                            <input
                                id="password"
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                required
                                placeholder="Password"
                                className="w-full px-4 py-3 rounded-xl bg-[#f0f2f4] text-base text-[#111418] placeholder-[#637588] focus:outline-none"
                            />
                        </div>
                    </div>
                    <button
                        type="submit"
                        disabled={loading}
                        className="group relative flex w-full justify-center rounded-xl bg-[#111418] py-3 px-4 text-sm font-semibold text-white hover:bg-black focus:outline-none"
                    >
                        {loading ? "Signing in..." : "Sign In"}
                    </button>
                </form>
                <p className="mt-4 text-center text-sm text-[#637588]">
                    Don't have an account? <a href="/signup" className="font-medium text-[#111418]">Sign up</a>
                </p>
            </div>
        </div>
    );
};
