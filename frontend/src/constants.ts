import { Duration } from "luxon";

export const formatHMSDuration = (seconds: number) => {
    return Duration.fromMillis(seconds * 1000).toFormat("hh:mm:ss")
};

export const formatHrDuration = (seconds?: number, fallback: string = "-") => {
    if (seconds === undefined || seconds === null) {
        return fallback;
    }
    
    return (seconds / 3600.0).toFixed(2) + "h";
}

export const formatReadibleDuration = (seconds?: number, fallback: string = "-") => {
    if (seconds === undefined || seconds === null) {
        return fallback;
    }

    const duration = Duration.fromMillis(seconds * 1000)
    const minutes = duration.as('minutes');
    const hours = duration.as('hours');
    
    if (hours < 2) {
        return minutes.toString() + 'm';
    }

    if (hours < 8) {
        return hours.toString() + 'h';
    }

    const workDays = hours / 8;

    return workDays.toString() + 'd';
}