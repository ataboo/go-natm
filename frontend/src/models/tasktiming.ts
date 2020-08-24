import moment from "moment";

export interface TaskTiming {
    current: moment.Duration
    estimated: moment.Duration|null
}