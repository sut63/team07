import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import ButtonGroup from '@material-ui/core/ButtonGroup';

import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';

const MainDriverPage: FC<{}> = () => {
const profile = { givenName: 'ข้อมูลการลงทะเบียนรถเข้าออก' };
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
    
    </Header>
     <Content>
       <ContentHeader title="">
       </ContentHeader>
       <br></br>
       <br></br>
       <br></br>
       <br></br>
       <center>
       <ButtonGroup
        orientation="vertical"
        color="primary"
        aria-label="vertical contained primary button group"
        variant="text"
      >
          <Link component={RouterLink} to="/createambulance">
        <Button><h1>ลงทะเบียนรถพยาบาล</h1></Button>
        </Link>
        <Link component={RouterLink} to="/Carcheckinout">
        <Button><h1>ลงทะเบียนรถเข้าออก</h1></Button>
        </Link>
        <Link component={RouterLink} to="/CreateTransport">
        <Button><h1>ส่งผู้ป่วย</h1></Button>
        </Link>
      </ButtonGroup>
       </center>

     </Content>
   </Page>
 );
};
 
export default MainDriverPage;