import React, { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// Fetch monitored sites
const fetchMonitoredSites = async () => {
    const res = await fetch("http://localhost:8080/status");
    if (!res.ok) throw new Error("Network response was not ok");
    return res.json();
};

// Post a new URL
const postNewSite = async (url: string) => {
    const res = await fetch("http://localhost:8080/urls", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ url }),
    });
    if (!res.ok) throw new Error("Failed to add new website");
    return res.json();
};

export const Dashboard: React.FC = () => {
    const queryClient = useQueryClient();
    const [newUrl, setNewUrl] = useState("");

    const { data, isLoading, error } = useQuery({
        queryKey: ['monitoredSites'],
        queryFn: fetchMonitoredSites
    });

    const mutation = useMutation({
        mutationFn: postNewSite,
        onSuccess: () => {
            setNewUrl("");  // Clear input
            queryClient.invalidateQueries({ queryKey: ['monitoredSites'] });  // Refetch data
        }
    });

    const handleAddSite = (e: any) => {
        e.preventDefault();
        if (newUrl.trim() === "") return;
        mutation.mutate(newUrl);
    };

    if (isLoading) return <div className="p-10 text-center">Loading...</div>;
    if (error) return <div className="p-10 text-center text-red-500">Error fetching data</div>;

    return (
        <div className="relative flex min-h-screen flex-col bg-white overflow-x-hidden font-[Manrope,_Noto_Sans,_sans-serif]">
            <div className="flex h-full grow flex-col">
                <header className="flex items-center justify-between border-b border-[#f0f2f4] px-10 py-3">
                    <div className="flex items-center gap-4 text-[#111418]">
                        <div className="w-4 h-4">
                            <svg viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M44 4H30.6666V17.3334H17.3334V30.6666H4V44H44V4Z" fill="currentColor" />
                            </svg>
                        </div>
                        <h2 className="text-lg font-bold tracking-[-0.015em]">SiteGuard</h2>
                    </div>
                    <div className="flex flex-1 justify-end gap-8">
                        <div className="flex items-center gap-9">
                            {['Overview', 'Docs', 'Community', 'Help'].map(link => (
                                <a key={link} href="#" className="text-sm font-medium text-[#111418]">{link}</a>
                            ))}
                        </div>
                        <div
                            className="bg-center bg-no-repeat bg-cover rounded-full w-10 h-10"
                            style={{ backgroundImage: 'url(https://cdn.usegalileo.ai/sdxl10/e12773e4-9685-4bc8-9099-0a8c04b4538d.png)' }}
                        />
                    </div>
                </header>

                <div className="px-40 flex flex-1 justify-center py-5">
                    <div className="flex flex-col max-w-[960px] flex-1">
                        <div className="flex flex-wrap justify-between gap-3 p-4">
                            <div className="flex min-w-72 flex-col gap-3">
                                <p className="text-[32px] font-bold text-[#111418]">Website Monitor</p>
                                <p className="text-sm text-[#637588]">Input the URL to inspect its health and key performance metrics.</p>
                            </div>
                        </div>

                        <div className="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
                            <label className="flex flex-col min-w-40 flex-1">
                                <form onSubmit={handleAddSite}>
                                    <input
                                        value={newUrl}
                                        onChange={(e) => setNewUrl(e.target.value)}
                                        placeholder="https://www.example.com"
                                        className="h-14 p-4 rounded-xl bg-[#f0f2f4] text-base text-[#111418] placeholder-[#637588] focus:outline-none w-full"
                                    />
                                </form>
                            </label>
                        </div>

                        {mutation.isError && (
                            <div className="text-red-500 px-4">Failed to add website. Try again.</div>
                        )}

                        <h3 className="text-lg font-bold text-[#111418] px-4 pb-2 pt-4">Monitored Sites</h3>
                        <div className="px-4 py-3">
                            <div className="flex overflow-hidden rounded-xl border border-[#dce0e5] bg-white">
                                <table className="flex-1">
                                    <thead>
                                        <tr className="bg-white">
                                            <th className="px-4 py-3 text-left text-sm font-medium text-[#111418] w-[400px]">URL</th>
                                            <th className="px-4 py-3 text-left text-sm font-medium text-[#111418] w-60">Status</th>
                                            <th className="px-4 py-3 text-left text-sm font-medium text-[#111418] w-[400px]">Response Time</th>
                                            <th className="px-4 py-3 text-left text-sm font-medium text-[#111418] w-[400px]">Availability</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        {data?.map((site: any) => (
                                            <tr key={site.url} className="border-t border-[#dce0e5]">
                                                <td className="px-4 py-2 w-[400px] text-sm text-[#111418]">{site.url}</td>
                                                <td className="px-4 py-2 w-60">
                                                    {site.is_up ? `✅` : `🔻`}
                                                </td>
                                                <td className="px-4 py-2 w-[400px] text-sm text-[#637588]">{site.response_time}ms</td>
                                                <td className="px-4 py-2 w-[400px] text-sm">
                                                    <div className="flex items-center gap-3">
                                                        <div className="w-[88px] overflow-hidden rounded-sm bg-[#dce0e5]">
                                                            <div className="h-1 rounded-full bg-[#111418]" style={{ width: `${site.availability}%` }}></div>
                                                        </div>
                                                        <p className="text-sm font-medium text-[#111418]">{site.availability}</p>
                                                    </div>
                                                </td>
                                            </tr>
                                        ))}
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div >
            </div >
        </div >
    );
};
