import { jwtDecode } from "jwt-decode";


export function getUser() {
    const token = localStorage.getItem("token")
    if (!token) {
        return ""
    }
    const decoded = jwtDecode(token);
    return decoded
}