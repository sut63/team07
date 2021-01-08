import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn';
import Ambulance from './components/Ambulancecreate';
import Ambulancetable from './components/listambulance';
import Mainambulance from './components/Ambulancemain';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/createambulance', Ambulance);
    router.registerRoute('/listambulance', Ambulancetable);
    router.registerRoute('/mainambulance', Mainambulance);
  },
});
