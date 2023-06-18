import { sveltekit } from '@sveltejs/kit/vite';

/** @type {import('vite').UserConfig} */
const config = {
  server: {
    fs: {
      // To allow serving files from the frontend project root:
      // allow: ['.'],
    },
  },
	plugins: [sveltekit()],
  resolve: {
    alias: {
      // This alias finishes the ability to reference our frontend directory
      // with "@path/to/file". You also need to add the path to jsconfig.json:
      // '@': path.resolve(__dirname, './'),
    },
  },
};

export default config;
