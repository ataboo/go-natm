import { Duration } from "luxon";

export const formatHMSDuration = (seconds: number) => {
    return Duration.fromMillis(seconds * 1000).toFormat("hh:mm:ss")
};

export const formatHrDuration = (seconds?: number, fallback: string = "-") => {
    if (!seconds) {
        return fallback;
    }
    
    return (seconds / 3600.0).toFixed(2) + "h";
}