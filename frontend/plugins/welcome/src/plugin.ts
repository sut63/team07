import { createPlugin } from '@backstage/core';
//import WelcomePage from './components/Transport';
import Ambulance from './components/Ambulancecreate';
import Ambulancetable from './components/listambulance';
import Mainambulance from './components/Ambulancemain';
import CarInspectionPage from './components/CarInspectionPage';
//import Carservicetable from './components/Carservicetable';
import Carservicecreate from './components/Carservicecreate';
import Carservicemain from './components/Carservicemain';
import Carcheckinout from './components/Carcheckinoutmain'
import CreateCarcheckinout from './components/Carcheckinoutcreate'
import CreateTransport from './components/Transport'
import LoginPage from './components/LoginPage'
import MainDriverPage from './components/Maindriver'
import CreateCarrepairrecord from './components/Carrepairrecordmain'
import Carrepairrecordtable from './components/CarrepairrecordTable';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', LoginPage);
    router.registerRoute('/createambulance', Ambulance);
    router.registerRoute('/listambulance', Ambulancetable);
    router.registerRoute('/mainambulance', Mainambulance);
    router.registerRoute('/carinspection', CarInspectionPage);
    router.registerRoute('/carservicemain', Carservicemain);
    router.registerRoute('/carservicecreate', Carservicecreate);
    router.registerRoute('/carcheckinout', Carcheckinout);
    router.registerRoute('/createcarcheckinout', CreateCarcheckinout);
    router.registerRoute('/createtransport', CreateTransport);
    router.registerRoute('/maindriver', MainDriverPage);
    router.registerRoute('/carrepairmain', CreateCarrepairrecord);
    router.registerRoute('/Carrepairrecordtable', Carrepairrecordtable);
  },
});
