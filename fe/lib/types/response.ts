export type Error = {
  success: false;
  code: number;
  message: string;
};

export type Success<T> = {
  success: true;
  code: number;
  message: string;
  data?: T;
};

export type ApiResponse<T = []> = Error | Success<T>;