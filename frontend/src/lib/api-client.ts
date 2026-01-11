import { OpenAPI } from "./api/core/OpenAPI";

// Configure the OpenAPI client
OpenAPI.BASE = import.meta.env.VITE_API_URL || "";
OpenAPI.WITH_CREDENTIALS = true;
OpenAPI.CREDENTIALS = "include";

// Re-export the AuthService and ApiError for easy access
export { AuthService } from "./api/services/AuthService";
export { ApiError } from "./api/core/ApiError";

// Re-export types
export type { dto_SuccessResponse } from "./api/models/dto_SuccessResponse";
export type { dto_ErrorResponse } from "./api/models/dto_ErrorResponse";
export type { dto_UserResponse } from "./api/models/dto_UserResponse";
export type { handlers_LoginRequest } from "./api/models/handlers_LoginRequest";
export type { handlers_SignupRequest } from "./api/models/handlers_SignupRequest";

/**
 * Example Usage:
 *
 * // Signup
 * const user = await AuthService.postAuthSignup({
 *   name: "John Doe",
 *   email: "john@example.com",
 *   password: "secure123",
 * });
 *
 * // Login
 * const result = await AuthService.postAuthLogin({
 *   email: "john@example.com",
 *   password: "secure123",
 * });
 *
 * // Get current user
 * const currentUser = await AuthService.getMe();
 *
 * // Logout
 * await AuthService.postAuthLogout();
 *
 * // Error handling
 * try {
 *   await AuthService.postAuthLogin({ email, password });
 * } catch (error) {
 *   if (error instanceof ApiError) {
 *     const errorBody = error.body as { error?: string };
 *     console.error(errorBody?.error);
 *   }
 * }
 */
