export enum LogLevel {
  INFO = 'info',
  WARN = 'warn',
  ERROR = 'error',
  DEBUG = 'debug',
  FATAL = 'fatal',
  UNKNOWN = 'unknown',
}

export interface LogEntry {
  level: string;
  time: string;
  message: string;
  fields?: Record<string, string>;
}

export interface LogFile {
  path: string;
  date: string;
  size: number;
}

export interface LogFilesResponse {
  files: LogFile[];
}

export interface PaginatedLogResponse {
  data: LogEntry[];
  total: number;
  page: number;
  limit: number;
}

export interface HealthInfo {
  uptime: string;
  uptimeSeconds: number;
  goRoutines: number;
  memAllocMb: number;
  memSysMb: number;
  numGc: number;
  dbConnected: boolean;
  goVersion: string;
  numCpu: number;
}
