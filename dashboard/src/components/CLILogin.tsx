import React, { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';

export const CLILogin: React.FC = () => {
    const [searchParams] = useSearchParams();
    const sessionId = searchParams.get('session');
    const [status, setStatus] = useState<'pending' | 'success' | 'error'>('pending');

    useEffect(() => {
        const token = localStorage.getItem('authToken');
        if (!token) {
            window.location.href = `/login?redirect=/cli-login?session=${sessionId}`;
            return;
        }

        fetch(`http://localhost:8080/cli/session/${sessionId}/complete`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ token }),
        })
            .then(res => {
                if (res.ok) {
                    setStatus('success');
                } else {
                    setStatus('error');
                }
            });
    }, [sessionId]);

    if (status === 'pending') {
        return <p>Authenticating CLI... Please wait.</p>;
    }

    if (status === 'success') {
        return (
            <div className="flex items-center justify-center min-h-screen">
                <div className="text-center">
                    <h1 className="text-2xl font-bold mb-4">✅ CLI Authenticated!</h1>
                    <p>You can now return to your terminal and close this tab.</p>
                </div>
            </div>
        );
    }

    return (
        <div className="flex items-center justify-center min-h-screen">
            <div className="text-center">
                <h1 className="text-2xl font-bold mb-4 text-red-500">⚠️ Authentication Failed</h1>
                <p>Please try running <code>uptime-monitor login</code> again.</p>
            </div>
        </div>
    );
};
