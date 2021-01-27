import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Carcheckinouttable';
import Button from '@material-ui/core/Button';
import HomeIcon from '@material-ui/icons/Home';
import CreateIcon from '@material-ui/icons/Create';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
const profile = { givenName: 'ข้อมูลการลงทะเบียนรถเข้าออก' };
const HeaderCustom = {
    minHeight: '50px',
  };
  
 return (
    
    <Page theme={pageTheme.other} >
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
     <Header
       style={HeaderCustom}
       title={`${profile.givenName}`}
       subtitle=""
     >    
     
        <div>
          <Link component={RouterLink} to="/maindriver">
            <Button variant="contained" color="primary" style={{backgroundColor: "#9834AE"}} startIcon={<HomeIcon/>}>
              หน้าแรก
            </Button>
          </Link> 
          <br></br>   
         </div>
    </Header>
     <Content>
       <ContentHeader title="ตารางการบันทึกรถเข้าออก">
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
       <Link component={RouterLink} to="/CreateCarcheckinout">
            <Button variant="contained" color="primary" style={{backgroundColor: "#9834AE"}} startIcon={<CreateIcon/>}>
              ลงทะเบียนรถเข้าออก
            </Button>
          </Link>
       </ContentHeader>
       <div>
       </div>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;