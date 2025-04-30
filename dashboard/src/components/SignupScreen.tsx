import { useState } from "react";

export const SignUpScreen: React.FC<{ onSignUp: (data: { firstName: string; lastName: string; email: string; password: string }) => void }> = ({ onSignUp }) => {
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    const handleSubmitSignUp = async (e: React.FormEvent) => {
        e.preventDefault();
        setError(null);

        // Basic validation
        if (password !== confirmPassword) {
            setError("Passwords do not match");
            return;
        }

        setLoading(true);
        try {
            await onSignUp({ firstName, lastName, email, password });
        } catch (err: any) {
            setError(err.message || "Sign up failed");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="flex min-h-screen items-center justify-center bg-white font-[Manrope,_Noto_Sans,_sans-serif]">
            <div className="w-full max-w-md space-y-8 p-10 rounded-2xl shadow-lg">
                <div className="text-center">
                    <h1 className="text-3xl font-bold text-[#111418]">Create an Account</h1>
                    <p className="mt-2 text-sm text-[#637588]">Join SiteGuard to monitor your sites</p>
                </div>
                {error && <div className="text-red-500 text-sm text-center">{error}</div>}
                <form className="mt-8 space-y-6" onSubmit={handleSubmitSignUp}>
                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label htmlFor="firstName" className="sr-only">First Name</label>
                            <input
                                id="firstName"
                                type="text"
                                value={firstName}
                                onChange={(e) => setFirstName(e.target.value)}
                                required
                                placeholder="First Name"
                                className="w-full px-4 py-3 rounded-xl bg-[#f0f2f4] text-base text-[#111418] placeholder-[#637588] focus:outline-none"
                            />
                        </div>
                        <div>
                            <label htmlFor="lastName" className="sr-only">Last Name</label>
                            <input
                                id="lastName"
                                type="text"
                                value={lastName}
                                onChange={(e) => setLastName(e.target.value)}
                                required
                                placeholder="Last Name"
                                className="w-full px-4 py-3 rounded-xl bg-[#f0f2f4] text-base text-[#111418] placeholder-[#637588] focus:outline-none"
                            />
                        </div>
                    </div>
                    <div>
                        <label htmlFor="email" className="sr-only">Email</label>
                        <input
                            id="email"
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            required
                            placeholder="Email address"
                            className="w-full px-4 py-3 rounded-xl bg-[#f0f2f4] text-base text-[#111418] placeholder-[#637588] focus:outline-none"
                        />
                    </div>
                    <div className="space-y-4">
                        <div>
                            <label htmlFor="signUpPassword" className="sr-only">Password</label>
                            <input
                                id="signUpPassword"
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                required
                                placeholder="Password"
                                className="w-full px-4 py-3 rounded-xl bg-[#f0f2f4] text-base text-[#111418] placeholder-[#637588] focus:outline-none"
                            />
                        </div>
                        <div>
                            <label htmlFor="confirmPassword" className="sr-only">Confirm Password</label>
                            <input
                                id="confirmPassword"
                                type="password"
                                value={confirmPassword}
                                onChange={(e) => setConfirmPassword(e.target.value)}
                                required
                                placeholder="Confirm Password"
                                className="w-full px-4 py-3 rounded-xl bg-[#f0f2f4] text-base text-[#111418] placeholder-[#637588] focus:outline-none"
                            />
                        </div>
                    </div>
                    <button
                        type="submit"
                        disabled={loading}
                        className="group relative flex w-full justify-center rounded-xl bg-[#111418] py-3 px-4 text-sm font-semibold text-white hover:bg-black focus:outline-none"
                    >
                        {loading ? "Creating account..." : "Sign Up"}
                    </button>
                </form>
                <p className="mt-4 text-center text-sm text-[#637588]">
                    Already have an account? <a href="/login" className="font-medium text-[#111418]">Sign in</a>
                </p>
            </div>
        </div>
    );
};