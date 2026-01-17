/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { dto_OAuthProviderResponse } from './dto_OAuthProviderResponse';
import type { dto_SocialAccountResponse } from './dto_SocialAccountResponse';
import type { models_Gender } from './models_Gender';
export type dto_UserResponse = {
    bio?: string;
    city?: string;
    createdAt?: string;
    date_of_birth?: string;
    email?: string;
    email_verified?: boolean;
    first_name?: string;
    gender?: models_Gender;
    id?: string;
    last_name?: string;
    oauth_providers?: Array<dto_OAuthProviderResponse>;
    phone?: string;
    phone_no_verified?: boolean;
    socials?: Array<dto_SocialAccountResponse>;
    updatedAt?: string;
    username?: string;
};

