import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Carservicetable';
import Button from '@material-ui/core/Button';
//import Typography from '@material-ui/core/Typography';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
const profile = { givenName: 'ระบบบันทึกการใช้รถพยาบาล' };

 return (
   <Page theme={pageTheme.library}>
     <Header
       title={`ท่านกำลังใช้งาน ${profile.givenName || ':)'}`}
       subtitle="สวัสดีครับท่านสมาชิกชมรมคนชอบรถพยาบาล"
     ><Link component={RouterLink} to="/Carservicesearch">
     <Button variant="contained" color="primary" >
       ไปยังหน้าค้นหา
     </Button>
   </Link>
   </Header>
     <Content>
       <ContentHeader title="ตารางบันทึกการใช้รถพยาบาล">
       
       <Link component={RouterLink} to="/Carservicecreate">
           <Button variant="contained" color="primary" >
             เพิ่มบันทึก
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
       <br></br>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;