/// <reference types="svelte" />
/// <reference types="vite/client" />
//
export declare global {
  interface Window {
    handleCaptchaCallback: (token: string) => Promise<void>;
    resetCaptcha: () => void;
    handleCaptchaError: () => void;
  }
}
