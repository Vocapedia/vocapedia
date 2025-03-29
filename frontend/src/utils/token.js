import { jwtDecode } from "jwt-decode";


export function getUser() {
    const token = localStorage.getItem("token")
    if (!token) {
        return ""
    }
    const decoded = jwtDecode(token);
    return decoded
}
export function getDevice() {

    const userAgent = navigator.userAgent;
    const platform = navigator.platform;
    const language = navigator.language;

    let brand = "Bilinmeyen", model = "Bilinmeyen", os = "Bilinmeyen";

    if (navigator.userAgentData) {
        const uaData = navigator.userAgentData;
        brand = uaData.brands?.[0]?.brand || "Bilinmeyen";
        model = uaData.mobile ? "Mobil Cihaz" : "Masaüstü";
        os = uaData.platform || "Bilinmeyen";
    } else {
        if (/Windows/i.test(userAgent)) os = "Windows";
        else if (/Mac OS/i.test(userAgent)) os = "MacOS";
        else if (/Linux/i.test(userAgent)) os = "Linux";
        else if (/Android/i.test(userAgent)) os = "Android";
        else if (/iPhone|iPad|iPod/i.test(userAgent)) os = "iOS";
    }

    let browser = "Bilinmeyen";
    if (/Edg/i.test(userAgent)) browser = "Microsoft Edge";
    else if (/Chrome/i.test(userAgent)) browser = "Google Chrome";
    else if (/Safari/i.test(userAgent)) browser = "Safari";
    else if (/Firefox/i.test(userAgent)) browser = "Mozilla Firefox";
    else if (/MSIE|Trident/i.test(userAgent)) browser = "Internet Explorer";


    return {
        brand,
        model,
        os,
        browser,
        platform,
        language,
    }
}