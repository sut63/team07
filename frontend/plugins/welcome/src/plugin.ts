import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Ambulance from './components/Ambulancecreate';
import Ambulancetable from './components/listambulance';
import Mainambulance from './components/Ambulancemain';
import CarInspectionPage from './components/CarInspectionPage';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/createambulance', Ambulance);
    router.registerRoute('/listambulance', Ambulancetable);
    router.registerRoute('/mainambulance', Mainambulance);
    router.registerRoute('/carinspection', CarInspectionPage);
  },
});
