import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async({ url, cookies }) => {
  const accessToken = url.searchParams.get('token') ?? ''
  cookies.set('mikuuuu', accessToken, { path: '/', httpOnly: true})

  throw redirect(302, '/')
}