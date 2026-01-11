/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { dto_SuccessResponse } from '../models/dto_SuccessResponse';
import type { dto_UserResponse } from '../models/dto_UserResponse';
import type { handlers_LoginRequest } from '../models/handlers_LoginRequest';
import type { handlers_SignupRequest } from '../models/handlers_SignupRequest';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class AuthService {
    /**
     * Login to user account
     * Authenticate user with email and password
     * @param request Login Request
     * @returns dto_SuccessResponse OK
     * @throws ApiError
     */
    public static postAuthLogin(
        request: handlers_LoginRequest,
    ): CancelablePromise<dto_SuccessResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/login',
            body: request,
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * Logout from user account
     * Invalidate user session and clear cookies
     * @returns dto_SuccessResponse OK
     * @throws ApiError
     */
    public static postAuthLogout(): CancelablePromise<dto_SuccessResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/logout',
        });
    }
    /**
     * Create a new user account
     * Register a new user with email and password
     * @param request Signup Request
     * @returns dto_UserResponse Created
     * @throws ApiError
     */
    public static postAuthSignup(
        request: handlers_SignupRequest,
    ): CancelablePromise<dto_UserResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/signup',
            body: request,
            errors: {
                400: `Bad Request`,
                409: `Conflict`,
                500: `Internal Server Error`,
            },
        });
    }
    /**
     * Get current user
     * Returns logged-in user
     * @returns dto_UserResponse OK
     * @throws ApiError
     */
    public static getUsersMe(): CancelablePromise<dto_UserResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/users/me',
            errors: {
                401: `Unauthorized`,
            },
        });
    }
}
