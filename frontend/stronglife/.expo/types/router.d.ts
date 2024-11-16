/* eslint-disable */
import * as Router from 'expo-router';

export * from 'expo-router';

declare module 'expo-router' {
  export namespace ExpoRouter {
    export interface __routes<T extends string = string> extends Record<string, unknown> {
      StaticRoutes: `/` | `/AppTabs` | `/Onboarding` | `/_sitemap` | `/sign_in`;
      DynamicRoutes: never;
      DynamicRouteTemplate: never;
    }
  }
}
