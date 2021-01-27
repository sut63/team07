import { createPlugin } from '@backstage/core';
//import WelcomePage from './components/Transport';
import Ambulance from './components/Ambulancecreate';
import Ambulancetable from './components/listambulance';
import Mainambulance from './components/Ambulancemain';
import CarInspectionPage from './components/CarInspectionPage';
import Carservicecreate from './components/Carservicecreate';
import Carservicemain from './components/Carservicemain';
import Carcheckinout from './components/Carcheckinoutmain'
import CreateCarcheckinout from './components/Carcheckinoutcreate'
import CreateTransport from './components/Transport'
import TransportTable from './components/TransportTable'
import SearchTransport from './components/SearchTransport'
import LoginPage from './components/LoginPage'
import MainDriverPage from './components/Maindriver'
import CreateCarrepairrecord from './components/Carrepairrecordmain'
import Carrepairrecordtable from './components/CarrepairrecordTable';
import Carservicesearch from './components/Carservicesearch';
import Searchambulance from './components/Searchambulance';
import CarInspectionSearchPage from './components/CarInspectionSearchPage';
import CarrepairrecordSearch from './components/Carrepairrecordsearch';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', LoginPage);
//CarInspection Link
    router.registerRoute('/carinspection', CarInspectionPage);
    router.registerRoute('/carinspectionsearch', CarInspectionSearchPage);
//Carrepair Link
    router.registerRoute('/carrepairmain', CreateCarrepairrecord);
    router.registerRoute('/Carrepairrecordtable', Carrepairrecordtable);
    router.registerRoute('/carrepairrecordsearch', CarrepairrecordSearch);
//Carservice Link
    router.registerRoute('/carservicemain', Carservicemain);
    router.registerRoute('/carservicecreate', Carservicecreate);
    router.registerRoute('/carservicesearch', Carservicesearch);
//MainDeiver
    router.registerRoute('/maindriver', MainDriverPage);
//Ambulance Link
    router.registerRoute('/createambulance', Ambulance);
    router.registerRoute('/listambulance', Ambulancetable);
    router.registerRoute('/searchambulance', Searchambulance);
    router.registerRoute('/mainambulance', Mainambulance);
//Carcheckinout Link     
    router.registerRoute('/carcheckinout', Carcheckinout);
    router.registerRoute('/createcarcheckinout', CreateCarcheckinout);
//Transport Link    
    router.registerRoute('/createtransport', CreateTransport);
    router.registerRoute('/transporttable', TransportTable);
    router.registerRoute('/searchtransport',SearchTransport);
    router.registerRoute('/carservicesearch', Carservicesearch);
    router.registerRoute('/searchambulance', Searchambulance);
    router.registerRoute('/carinspectionsearch', CarInspectionSearchPage);
  },
});
