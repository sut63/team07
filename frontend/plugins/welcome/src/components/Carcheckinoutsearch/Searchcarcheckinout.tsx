import { useState, useEffect } from 'react';
import * as React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { fade, makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import Avatar from '@material-ui/core/Avatar';
import { anfaBase64Function } from '../../image/anfa';
import { EntUser } from '../../api/models/EntUser';
import { EntCarCheckInOut } from '../../api/models/EntCarCheckInOut';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import TextField from '@material-ui/core/TextField';
import { Alert } from '@material-ui/lab';
// import SearchIcon from "@material-ui/icons/Search";

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
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    formControl: {
      width: 400,
    },
    table: {
        minWidth: 1500,
      },
      search: {
        position: 'relative',
      borderRadius: theme.shape.borderRadius,
      backgroundColor: fade(theme.palette.common.white, 0.15),
      '&:hover': {
        backgroundColor: fade(theme.palette.common.white, 0.25),
      },
      marginLeft: 0,
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing(1),
        width: '20%',
        }
      },
      searchIcon: {
        padding: theme.spacing(0, 2),
      height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      },
    inputRoot: {
        color: "inherit"
      },
      inputInput: {
        padding: theme.spacing(1, 1, 1, 0),
      // vertical padding + font size from searchIcon
      paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
      transition: theme.transitions.create('width'),
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        width: '12ch',
        '&:focus': {
          width: '20ch',
          }
        }
      }
  }),
);

export default function Searchcarcheckinout() {
  const classes = useStyles();
  const profile = { givenName: 'ระบบลงทะเบียนรถเข้าออก' };
  const api = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [loading, setLoading] = useState(true);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  const [search, setSearch] = useState(String);

  const [userid, setUser] = useState(Number);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [carcheckinout, setCarcheckinout] = useState<EntCarCheckInOut[]>([]);

  useEffect(() => {

    const getUsers = async () => {
        const us = await api.listUser();
          setLoading(false);
          setUsers(us);
        };
        getUsers();

    const getCarcheckinouts = async () => {
        const res = await api.listCarcheckinout({ limit: 120, offset: 0 });
        setLoading(false);
        setCarcheckinout(res);
      };
      getCarcheckinouts();

    const checkJobPosition = async () => {
        const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
        setLoading(false);
        if (jobdata != "เจ้าหน้าที่รถพยาบาล" ) {
          localStorage.setItem("userdata",JSON.stringify(null));
          localStorage.setItem("jobpositiondata",JSON.stringify(null));
          history.pushState("","","./");
          window.location.reload(false);        
        }
        else{
            setUser(Number(localStorage.getItem("userdata")))
        }
      }
    checkJobPosition();
 
  }, [loading]);  

  const SearchhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(event.target.value as string);
  };

  const UserhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

  const Searchcarinout = (res : any) =>{
    const apiUrl = `http://localhost:8080/api/v1/carcheckinoutsearch?ambulance=${search}`;
    const requestOptions = {
      method: 'GET',
  };  
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
        .then(data => {
          console.log(data.data)
          setErrorMessege("ไม่พบข้อมูล");
          setAlertType("error");
          setCarcheckinout([]);
          if (data.data != null) {
            if(data.data.length >= 1) {
              setErrorMessege("พพบข้อมูล");
              setAlertType("success");
              console.log(data.data)
              setCarcheckinout(data.data);
          }
      }

      setStatus(true);
  });

}

  return (
 <Page theme={pageTheme.other}>
      <Header
        title={`${profile.givenName}`}
      >
        <Button
          style={{ marginLeft: 20 ,backgroundColor: "#9834AE"}}
          component={RouterLink}
          to="/Carcheckinout"
          variant="contained"
          color="primary"
        >
          กลับ
       </Button>
      </Header>
      <Content>
        <ContentHeader title="ค้นหารถเข้าออก">
        <Avatar alt="anfa" src={anfaBase64Function} /> &nbsp;&nbsp;&nbsp;&nbsp;
        <h3>{users.filter((filter: EntUser) => filter.id == userid).map((item: EntUser) => `${item.name} (${item.email})`)}</h3>
        &nbsp;&nbsp;&nbsp;&nbsp;
        {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
        </ContentHeader>   
        <div className={classes.root}>
        <form noValidate autoComplete="off">
            <div className={classes.margin}>
              <div>
            <TextField
                 id="search"
                 label = "ค้นหาทะเบียนรถ"
                //  variant="standard"
                //  color="secondary"
                 type="string"
                 value={search}
                 onChange={SearchhandleChange}
                 className={classes.textField}
                 style={{ width: 200 }}
                />
            &nbsp;&nbsp;&nbsp;&nbsp;
            <Button
                onClick={() => {
                  Searchcarinout();
                }}
                variant="contained"
                color="primary"
              >
                ค้นหา
             </Button>
             </div>
          </div>
          <br></br><br></br>

        <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ทะเบียนรถ</TableCell>
           <TableCell align="center">เจ้าหน้าที่</TableCell>
           <TableCell align="center">วัตถุประสงค์</TableCell>
           <TableCell align="center">หมายเหตุ</TableCell>
           <TableCell align="center">จำนวนเจ้าหน้าที่</TableCell>
           <TableCell align="center">ระยะทาง</TableCell>
           <TableCell align="center">สถานที่</TableCell>
           <TableCell align="center">วันที่รถออก</TableCell>
           <TableCell align="center">วันที่รถเข้า</TableCell>
           
         </TableRow>
       </TableHead>
       <TableBody>
         {carcheckinout.map((item:any )=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.edges.ambulance.carregistration}</TableCell>
             <TableCell align="center">{item.edges.name.name}</TableCell>
             <TableCell align="center">{item.edges.purpose.objective}</TableCell>
             <TableCell align="center">{item.note}</TableCell>
             <TableCell align="center">{item.person}</TableCell>
             <TableCell align="center">{item.distance}</TableCell>
             <TableCell align="center">{item.place}</TableCell>
             <TableCell align="center">{moment(item.checkOut).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">{moment(item.checkIn).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>            
          </form>
        </div>
      </Content>
    </Page>
  );
}
