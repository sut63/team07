import React, { FC, useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis';
import Button from '@material-ui/core/Button';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import { jaikeleBase64Function } from '../../image/jaikele'
import CardMedia from '@material-ui/core/CardMedia';
import FormControl from '@material-ui/core/FormControl';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import AmbulancecarIcon from '@material-ui/icons/AirportShuttle';
import ListAltIcon from '@material-ui/icons/ListAlt';
import SearchIcon from "@material-ui/icons/Search";
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link
} from '@backstage/core';
import { EntUser } from '../../api/models/EntUser';

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
const api = new DefaultApi();
const classes = useStyles();  
const [users, setUsers] = useState<EntUser[]>([]);
const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getUsers = async () => {
            const res = await api.listUser();
            setLoading(false);
            setUsers(res);
        };
        getUsers();

        const checkJobPosition = async () => {
          const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
          setLoading(false);
          if (jobdata != "เจ้าหน้าที่รถพยาบาล" ) {
            localStorage.setItem("userdata",JSON.stringify(null));
            localStorage.setItem("jobpositiondata",JSON.stringify(null));
            history.pushState("","","./");
            window.location.reload(false);        
          }
        }
      checkJobPosition();
   
    }, [loading]);

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
        <Button startIcon={<AmbulancecarIcon/>}><h1>ลงทะเบียนรถพยาบาล</h1></Button>
        </Link>
        <Link component={RouterLink} to="/Carcheckinout">
        <Button startIcon={<ListAltIcon/>}><h1>ลงทะเบียนรถเข้าออก</h1></Button>
        </Link>
        <Link component={RouterLink} to="/CreateTransport">
        <Button startIcon={<ShoppingCartIcon/>}><h1>ส่งผู้ป่วย</h1></Button>
        </Link>
        <Link component={RouterLink} to="/Searchambulance">
        <Button startIcon={<SearchIcon/>}><h1>ค้นหารถพยาบาล</h1></Button>
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
          loop
          muted
       />
       </FormControl>
       </center>

     </Content>
   </Page>
 );
};
 
export default MainDriverPage;