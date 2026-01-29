export default function GetLocalIp(): string {
    const localIp = window.location.hostname

    return localIp;
}