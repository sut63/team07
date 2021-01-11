import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import { jaikeleBase64Function } from '../../image/jaikele'
import CardMedia from '@material-ui/core/CardMedia';
import FormControl from '@material-ui/core/FormControl';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link
} from '@backstage/core';

const MainDriverPage: FC<{}> = () => {
const profile = { givenName: 'ระบบรถพยาบาล' };
const HeaderCustom = {
    minHeight: '50px',
  };
  const useStyles = makeStyles((theme: Theme) =>
  createStyles({
      root: {
          display: 'flex',
          flexWrap: 'wrap',
          justifyContent: 'center',
      },
      margin: {
          margin: theme.spacing(3),
      },
      withoutLabel: {
          marginTop: theme.spacing(3),
      },
      textField: {
          width: '25ch',
      },
      media: {
          height: 0,
          marginLeft: 25,
          maxWidth: 300,
      }
      
  }),
);
const classes = useStyles();  
 return (
    <Page theme={pageTheme.other} >
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
       <center>
       <br></br>
       <br></br>
       
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
      <br></br>
      <FormControl
        className={classes.media}>
       <CardMedia
          component="video"
          src={jaikeleBase64Function}
          style={{ marginTop: 30 }}
          autoPlay
       />
       </FormControl>
       </center>

     </Content>
   </Page>
 );
};
 
export default MainDriverPage;