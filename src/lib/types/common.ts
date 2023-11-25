export interface ClientServiceResponse<T> {
  code: string;
  message: string;
  data: T
}

export interface User {
  userId: string;
}