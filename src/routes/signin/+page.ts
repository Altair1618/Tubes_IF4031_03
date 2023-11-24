import type { PageLoad } from "./$types";
import { PUBLIC_CLIENT_SERVICE_BASE_URL} from '$env/static/public'

export const load: PageLoad = () => {
  return {
    clientURL: PUBLIC_CLIENT_SERVICE_BASE_URL
  }
}