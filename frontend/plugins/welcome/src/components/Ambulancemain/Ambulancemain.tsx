import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../listambulance';
import Button from '@material-ui/core/Button';


import TableRow from '@material-ui/core/TableRow';

 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
import { TableCell, TableContainer } from '@material-ui/core';
 
const WelcomePage: FC<{}> = () => {
const profile = { givenName: 'ข้อมูลการลงทะเบียนรถ' };
const HeaderCustom = {
    minHeight: '50px',
  };
 return (
    
    <Page theme={pageTheme.service} >
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
     <Header
       style={HeaderCustom}
       title={`${profile.givenName}`}
       subtitle=""
     >    
     
        
      <table>
        <tr>
          <th>
          <Link component={RouterLink} to="/createambulance">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              เพิ่มข้อมูลรถเข้าสู่ระบบ
            </Button>
          </Link>
          </th>
          <th>
          <Link component={RouterLink} to="/maindriver">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              กลับหน้าหลัก
            </Button>
          </Link>
          </th>
          </tr>
          </table>
    </Header>
     <Content>
       <ContentHeader title="รถพยาบาลในระบบ">
       </ContentHeader>
       <div>
       </div>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;