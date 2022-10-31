export enum ButtonVariant {
  PRIMARY = "primary",
  DANGER = "danger",
  WHITE = "white",
}

export enum ButtonTextSize {
  SMALL = "sm",
  MEDIUM = "md",
  LARGE = "lg",
}

export enum ToastType {
  INFO = "info",
  SUCCESS = "success",
  ERROR = "error",
}

export interface IToast {
  message: string;
  type: ToastType;
}
