/** @type {import('tailwindcss').Config} */

export default {
	content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
	theme: {
		colors: {
			gray: {
				'50': '#FFFFFF',
				'55': '#F5F5F5',
				'100': '#EFEFEF',
				'200': '#DCDCDC',
				'300': '#BDBDBD',
				'400': '#989898',
				'500': '#7C7C7C',
				'600': '#656565',
				'700': '#525252',
				'800': '#464646',
				'900': '#3D3D3D',
				'950': '#292929',
			},
			blue: {
				'50': '#F0F8FF',
				'100': '#DFF0FF',
				'200': '#CAE9FF',
				'300': '#79CAFF',
				'400': '#32B0FE',
				'500': '#0796F0',
				'600': '#0077CD',
				'700': '#005EA6',
				'800': '#035089',
				'900': '#094471',
				'950': '#062A4B',
			},
			red: {
				'50': '#FFF1F2',
				'100': '#FFE1E3',
				'200': '#FFC7CB',
				'300': '#FFA0A6',
				'400': '#FF5964',
				'500': '#F83B48',
				'600': '#E51D2A',
				'700': '#C1141F',
				'800': '#A0141D',
				'900': '#84181F',
				'950': '#48070B',
			},
			green: {
				'50': '#F2FDF0',
				'100': '#E1FCDC',
				'200': '#C3F7BB',
				'300': '#93EF86',
				'400': '#5BDF49',
				'500': '#34C421',
				'600': '#25A415',
				'700': '#218015',
				'800': '#1E6516',
				'900': '#1A5314',
				'950': '#082E05',
			},
			black: "#000000",
			white: "#FFFFFF",
		},
		fontFamily: {
			kanit: ['Kanit', 'sans-serif']
		},
		extend: {}
	},
	plugins: [],
}
