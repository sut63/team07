import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { EntInspectionResult } from '../../api/models/EntInspectionResult';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import { EntAmbulance } from '../../api/models/EntAmbulance';
import Swal from 'sweetalert2'
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
import { Page, pageTheme, Header, Content, Link} from '@backstage/core';
import { Grid, MenuItem, Button, TextField, Select, Typography, FormControl, InputLabel} from '@material-ui/core';
import { EntUser } from '../../api/models/EntUser';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    headsearch:{
     width:'auto',
     margin:'10px',
     color:'#FFFFFF',
     background: '#2196F3',
    },
    margin: {
      margin: theme.spacing(1),
    },
    margins: {
      margin: theme.spacing(2),
    },
    
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(3),
      marginBottom: theme.spacing(3),
    },
    table: {
      minWidth: 500,
    },
  
    
  }),
);
const Toast = Swal.mixin({
 
  position: 'center',
  showConfirmButton: false,
   showCloseButton: true,
  
});


export default function ComponentsTable() {
const [carstatuses, setCarstatuses] = useState<EntInspectionResult[]>([]);
const [ambulances, setAmbulances] = useState<EntAmbulance[]>([]);
 const [carstatusid, setcarstatus] = useState(Number);
 const classes = useStyles();
 const api = new DefaultApi();
 const [loading, setLoading] = useState(true);
 const [search, setSearch] = useState(false);
 const [userid, setUser] = useState(Number);
 const [registration, setregistration] = useState(String);
 const [users, setUsers] = useState<EntUser[]>([]);
 const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบค้นหาข้อมูลรถ' };

 const alertMessage = (icon: any, title: any) => {
  Toast.fire({
    icon: icon,
    title: title,
  });
  setSearch(false);
}

 useEffect(() => {
   const getCarstatuses = async () => {
 
    const ins = await api.listInspectionresult();
      setLoading(false);
      setCarstatuses(ins);
    };
    getCarstatuses();
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
  const getUsers = async () => {
 
    const us = await api.listUser();
      setLoading(false);
      setUsers(us);
    };
    getUsers();
checkJobPosition();
 }, [loading]);



 const statushandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  
  setcarstatus(event.target.value as number);
};
const Registrationhandlehange = (event: React.ChangeEvent<{  value: unknown }>) => {
 
  setregistration(event.target.value as string);
  
};

const cleardata = () =>{
  setcarstatus(0);
  setregistration("");
  
}

const SearchCarInspection = async () => {
  console.log(registration);
  if(carstatusid == 0 && registration == ""){
    alertMessage("info","แสดงข้อมูลรถทั้งหมดในระบบ");
    const res = await api.listAmbulance();
            setSearch(true);
            setAmbulances(res);
   }
   else{
  const apiUrl = `http://localhost:8080/api/v1/searchambulances?ambulance=${registration}&status=${carstatusid}`;
  const requestOptions = {
      method: 'GET',
  };
  fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
          console.log(data.data)
          alertMessage("warning","ไม่พบข้อมูลที่ค้นหา")
          setAmbulances([]);
          if (data.data != null) {
              if(data.data.length >= 1) {
                alertMessage("success","พบข้อมูลที่ค้นหา")
                  console.log(data.data)
                  setAmbulances(data.data);
              }
          }

      
      });
    }
}



 return (
   
 <Page theme={pageTheme.home}>
   <Header
        title={`${profile.givenName}`}
      >
       <table>
       <tr>
         <th>
           <h3 style={
             {color: "#483D8B",
             background: 'linear-gradient(45deg, #FFFACD 15%, #9932CC 120%)',
             borderRadius: 10,
             height: 20,
             padding: '0 30px',
             
            }
          }>
            {users.filter((filter: EntUser) => filter.id == userid).map((item: EntUser) => `${item.name} (${item.email})`)}
        </h3>
        </th>
        <th>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          <Link component={RouterLink} to="/maindriver">
            <Button variant="contained"style={{ background: 'linear-gradient(45deg, #FFFACD 15%, #9932CC 120%)',height: 40}}>
              <h3
              style={
                {color: "#483D8B",
                background: 'linear-gradient(45deg, #FFFACD 15%, #9932CC 120%)',
                borderRadius: 10,
                height: 25,
                padding: '0 20px',
                
               }
             }>
               กลับหน้าหลัก
            </h3>
            </Button>
          </Link>
          </th>
       </tr>
       </table>

      </Header>
   <Content>
    <Grid container item xs={12} justify="center">
    <Grid item xs={5}>
          <Paper>
          
            
            <Typography align="center" > 
            <div  style={{ background: "#734C8F",height: 55}}>
            <h1 style={
             {color: "#FFFFFF",
             borderRadius: 10,
             height: 20,
             padding: '0 30px',
             fontSize:'40px',
            }}>
              ค้นหาข้อมูลรถ
            </h1>
            </div>

            <div>
            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                               <div className={classes.paper}><strong>เลขทะเบียนรถ</strong></div>
                                <TextField
                                    id="registration"
                                    value={registration}
                                    onChange={Registrationhandlehange}
                                    type="string"
                                    size="small"
                                   
                                    style={{ width: 200 }}
                                />
             </FormControl>
            
            
            </div>
           <div>
           <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>สถานะของรถ</strong></div>
              <InputLabel id="company-label"></InputLabel>
              <Select
                labelId="company-label"
                id="company"
                //value={insuranceid}
                //onChange={InsurancehandleChange}
                value={carstatusid}
                onChange={statushandleChange}
                style={{ width: 300 }}
                
              >
                <MenuItem value={0}>-</MenuItem>
                {carstatuses.filter((filter: any) => filter.edges.jobposition.positionName == "เจ้าหน้าที่รถพยาบาล").map((item: EntInspectionResult) => (
                    <MenuItem value={item.id}>{item.resultName}</MenuItem>
                   ))}
              </Select>
            </FormControl>
            </div>
            <div></div>
            <Button
             onClick={() => {
              SearchCarInspection();
              
            }}  
            endIcon={<SearchTwoToneIcon />}
            className={classes.margins} 
            variant="contained"
            style={{ background: "#734C8F",height: 40}}>
              <h3
              style={
                {color: "#FFFFFF",
                padding: '0 10px',
                
                }
             }>
               ค้นหาข้อมูล
            </h3>
            </Button>
            <Button
             onClick={() => {
              cleardata();
              
            }} 
            className={classes.margins} 
            variant="contained"
            style={{ background: "#734C8F",height: 40}}>
              <h3
              style={
                {color: "#FFFFFF",
                padding: '0 25px',
                
                }
             }>
               เคลียร์ข้อมูล
            </h3>
            </Button>
      
            </Typography>
            
                      
                        
           
          </Paper>
        </Grid>
    </Grid>


    <Grid container  justify="center">
     <Grid item xs={12} md={10}>
      <Paper>
        {search ?(
      <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                                <TableCell  align="center">ลำดับ</TableCell>
                                <TableCell align="center">ทะเบียนรถ</TableCell>
                                <TableCell align="center">แบรนด์รถ</TableCell>
                                <TableCell align="center">เครื่องยนต์</TableCell>
                                <TableCell align="center">ความจุตัวถัง</TableCell>
                                <TableCell align="center">บริษัทประกัน</TableCell>
                                <TableCell align="center">สถานะ</TableCell>
                                <TableCell align="center">เจ้าหน้าที่</TableCell>
                                <TableCell align="center">วันที่</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>
                     
                              {ambulances.map((item: any)=> (
                                <TableRow key={item.id}>
                                  <TableCell align="center">{item.id}</TableCell>
                                  <TableCell align="center">{item.carregistration}</TableCell>
                                  <TableCell align="center">{item.edges?.Hasbrand?.brand}</TableCell>
                                  <TableCell align="center">{item.enginepower}</TableCell>
                                  <TableCell align="center">{item.displacement}</TableCell>
                                  <TableCell align="center">{item.edges?.Hasinsurance?.company}</TableCell>
                                  <TableCell align="center">{item.edges.Hasstatus.resultName}</TableCell>
                                  <TableCell align="center">{item.edges?.Hasuser?.name}</TableCell>
                                  <TableCell align="center">{moment(item.registerat).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                                  
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>
        ):<TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell  align="center">ลำดับ</TableCell>
              <TableCell align="center">ทะเบียนรถ</TableCell>
              <TableCell align="center">แบรนด์รถ</TableCell>
              <TableCell align="center">เครื่องยนต์</TableCell>
              <TableCell align="center">ความจุตัวถัง</TableCell>
              <TableCell align="center">บริษัทประกัน</TableCell>
              <TableCell align="center">สถานะ</TableCell>
              <TableCell align="center">เจ้าหน้าที่</TableCell>
              <TableCell align="center">วันที่</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
   
            {ambulances.map((item: any)=> (
              <TableRow key={item.id}>
                <TableCell align="center">{item.id}</TableCell>
                <TableCell align="center">{item.carregistration}</TableCell>
                <TableCell align="center">{item.edges?.Hasbrand?.brand}</TableCell>
                <TableCell align="center">{item.enginepower}</TableCell>
                <TableCell align="center">{item.displacement}</TableCell>
                <TableCell align="center">{item.edges?.Hasinsurance?.company}</TableCell>
                <TableCell align="center">{item.edges.Hasstatus.result_name}</TableCell>
                <TableCell align="center">{item.edges?.Hasuser?.name}</TableCell>
                <TableCell align="center">{moment(item.registerat).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>}
      </Paper>
     </Grid>
     </Grid>
   </Content>
 </Page>
 );

}
