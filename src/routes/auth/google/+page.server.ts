import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async({ url, cookies }) => {
  const accessToken = url.searchParams.get('token') ?? ''
  console.log(accessToken)
  cookies.set('mikuuuu', accessToken, { path: '/', httpOnly: true})


  throw redirect(302, '/')
}