/* eslint-disable */
import * as Router from 'expo-router';

export * from 'expo-router';

declare module 'expo-router' {
  export namespace ExpoRouter {
    export interface __routes<T extends string = string> extends Record<string, unknown> {
      StaticRoutes: `/` | `/App` | `/_sitemap` | `/components/BotonesSelector` | `/components/Calendar` | `/components/Formulario` | `/components/Header` | `/components/LittleComponents/TabBar` | `/hooks/useColorScheme` | `/hooks/useThemeColor` | `/screens/Login` | `/screens/MiCocina` | `/screens/MiGimnasio` | `/screens/Planificacion` | `/screens/Training` | `/styles/Appstyles` | `/types`;
      DynamicRoutes: never;
      DynamicRouteTemplate: never;
    }
  }
}
