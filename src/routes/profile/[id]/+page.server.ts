import { PUBLIC_CLIENT_SERVICE_BASE_URL } from '$env/static/public';
import { error, fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import type { ClientServiceResponse } from '$lib/types/common';
import { message, setError, superValidate } from 'sveltekit-superforms/server';
import { updateProfileSchema } from '$lib/dto/profile/updateProfile.dto';

interface UserData {
	email: string;
	picture: string;
	name: string;
	id: string;
}

export const load: PageServerLoad = async ({ params, cookies }) => {
	const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/profile`, {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${cookies.get('mikuuuu')}`
		},
		credentials: 'include'
	});
	const responseData: ClientServiceResponse<UserData> = await response.json();

	if (!response.ok) {
		throw error(404, 'Not found');
	}

	const form = await superValidate({ name: responseData.data.name }, updateProfileSchema);

	return {
		form,
		...responseData
	};
};

export const actions = {
	updateProfile: async ({ request, cookies }) => {
		const form = await superValidate(request, updateProfileSchema);
		console.log('post', form);

		if (!form.valid) {
			return fail(400, { form });
		}

		try {
			console.log(form.data.name);
			const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/profile`, {
				method: 'PATCH',
				body: JSON.stringify({
					name: form.data.name
				}),
				headers: {
					Authorization: `Bearer ${cookies.get('mikuuuu')}`,
					'Content-Type': 'application/json'
				},
				credentials: 'include'
			});
			const responseData: ClientServiceResponse<{}> = await response.json();

			if (!response.ok) {
				return message(form, responseData.message, { status: 400 });
			}

			return {
				form,
				message: responseData.message
			};
		} catch (e) {
			console.log(e);
			return message(form, 'Something went wrong, please try again later', { status: 500 });
		}
	},
	signOut: async (event) => {
		try {
			const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/auth/signout`, {
				method: 'post',
				headers: {
					'Authorization': `Bearer ${event.cookies.get('mikuuuu')}`
				},
				credentials: 'include'
			})
			const responseData: ClientServiceResponse<{}> = await response.json()

			if (!response.ok) {
				return fail(500, { message: "Something went wrong, please try again later" })
			}

			event.cookies.delete('mikuuuu', { path: '/'})

			return {
				message: responseData.message
			}
		} catch (e) {
			console.log(e)
			return fail(500, { message: "Something went wrong, please try again later" })
		}
	},
	deleteProfile: async (event) => {
		try {
			const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/profile`, {
				method: "DELETE",
				headers: {
					'Authorization': `Bearer ${event.cookies.get('mikuuuu')}`
				},
				credentials: 'include'
			})

			const responseData: ClientServiceResponse<{}> = await response.json()


			if (!response.ok) {
				return fail(500, { message: "Something went wrong, please try again later" })
			}

			event.cookies.delete('mikuuuu', { path: '/'})

			return {
				message: responseData.message
			}

		} catch (e) {
			console.log(e)
			return fail(500, { message: "Something went wrong, please try again later" })
		}
	}
} satisfies Actions;
