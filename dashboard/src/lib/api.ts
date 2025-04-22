import axios from "axios";

export async function getStatuses() {
    const res = await axios.get("http://localhost:8080/status");
    return res.data;
}
