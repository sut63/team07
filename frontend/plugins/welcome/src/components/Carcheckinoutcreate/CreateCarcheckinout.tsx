import { useState, useEffect } from 'react';
import * as React from 'react';
import { Link as RouterLink } from 'react-router-dom';
// import SaveIcon from '@material-ui/icons/Save'; // icon save
// import Swal from 'sweetalert2'; // alert
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import Avatar from '@material-ui/core/Avatar';
import { anfaBase64Function } from '../../image/anfa';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
// import Typography from '@material-ui/core/Typography';
import { EntAmbulance } from '../../api/models/EntAmbulance';
import { EntUser } from '../../api/models/EntUser';
import { EntPurpose } from '../../api/models/EntPurpose';

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
  }),
);

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ระบบลงทะเบียนรถเข้าออก' };
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  //เก็บข้อมูลที่จะดึงมา
  const [ambulances, setAmbulances] = useState<EntAmbulance[]>([]);
  const [purposes, setPurposes] = useState<EntPurpose[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);

  const [loading, setLoading] = useState(true);
  const [checkin, setCheckin] = useState(String);
  const [checkout, setCheckout] = useState(String);

  const [ambulanceid, setambulance] = useState(Number);
  const [purposeid, setpurpose] = useState(Number);
  const [userid, setUser] = useState(Number);
  const [note, setnote] = useState(String);

  useEffect(() => {

    const getAmbulances = async () => {
 
      const a = await api.listAmbulance({ limit: 10, offset: 0 });
      setLoading(false);
      setAmbulances(a);
    };
    getAmbulances();
 
    const getUsers = async () => {
 
    const us = await api.listUser();
      setLoading(false);
      setUsers(us);
    };
    getUsers();
 
    const getPurposes = async () => {
 
     const pp = await api.listPurpose();
       setLoading(false);
       setPurposes(pp);
     };
     getPurposes();

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
  console.log(userid)
  const CheckinhandleChange = (event: any) => {
    setCheckin(event.target.value as string);
  };
  const CheckouthandleChange = (event: any) => {
    setCheckout(event.target.value as string);
  };

  const AmbulancehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setambulance(event.target.value as number);
  };

  const PurposehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setpurpose(event.target.value as number);
  };

  const UserhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };
  const Notehandlehange = (event: any) => {
    setnote(event.target.value as string);
  };

  const CreateCarcheckinout = async () => {
    if ((ambulanceid != null) && (purposeid != null) && (checkin != "") && (checkout != "") && (checkin != null) && (checkout != null)) {
    const carcheckinouts = {
      ambulance: ambulanceid,
      name     : userid,
      purpose  : purposeid,
      note     : note,
      checkin  : checkin + ":00+07:00",
      checkout : checkout + ":00+07:00"
    };
    
    console.log(carcheckinouts);
    const res: any = await api.createCarcheckinout({ carcheckinout: carcheckinouts });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      //window.location.reload(false);
    } 
  }
  else {
      setStatus(true)
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
      }, 10000);
  };
  
  return (
 <Page theme={pageTheme.other}>
      <Header
        title={`${profile.givenName}`}
      //subtitle="Some quick intro and links."
      ></Header>
      <Content>
        <ContentHeader title="ลงทะเบียนรถเข้าออก">
        <Avatar alt="anfa" src={anfaBase64Function} /> &nbsp;&nbsp;&nbsp;&nbsp;
        <h3>{users.filter((filter: EntUser) => filter.id == userid).map((item: EntUser) => `${item.name} (${item.email})`)}</h3>
        &nbsp;&nbsp;&nbsp;&nbsp;
          {status ? (
            <div>
              {alert ? (
                 <Alert severity="success" style={{ marginTop: 20 }} onClose={() => {setStatus(false)}}>
                  บันทึกสำเร็จ
                 </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }} onClose={() => {setStatus(false)}}>
                    กรุณากรอกข้อมูลอีกครั้ง
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <FormControl>
       
        </FormControl>
        <div className={classes.root}>
          <form noValidate autoComplete="off">

          <div>
          {/* <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>เจ้าหน้าที่รถพยาบาล</strong></div>
              <InputLabel id="staff-label"></InputLabel>
              <Select
                labelId="staff-label"
                id="user"
                value={userid}
                onChange={UserhandleChange}
                style={{ width: 400 }}
              >
                {users.map((item: EntUser) => (
                  <MenuItem value={item.id}>{item.name}</MenuItem>
                ))}
              </Select>
            </FormControl> */}
          <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>รถพยาบาล</strong></div>
              <InputLabel id="brand-label"></InputLabel>
              <Select
                labelId="ambulance-label"
                id="ambulance"
                value={ambulanceid}
                onChange={AmbulancehandleChange}
                style={{ width: 400 }}
              >
                {ambulances.map((item: EntAmbulance) => (
                  <MenuItem value={item.id}>{item.carregistration}</MenuItem>
                ))}
              </Select>
            </FormControl>
            </div>

          <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>วัตถุประสงค์</strong></div>
              <InputLabel id="purpose-label"></InputLabel>
              <Select
                labelId="purpose-label"
                id="purpose"
                value={purposeid}
                onChange={PurposehandleChange}
                style={{ width: 400 }}
              >
                {purposes.map((item: EntPurpose) => (
                  <MenuItem value={item.id}>{item.objective}</MenuItem>
                ))}
              </Select>
            </FormControl>
            </div>

          <div className={classes.paper}><strong>หมายเหตุ</strong></div>
            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}      
              id="note"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={note}
              onChange={Notehandlehange}
            />
       <br></br><br></br>
            <div className={classes.paper}><strong>วันที่เวลารถออก</strong></div>
      <center>             <FormControl className={classes.margin} >
        <div>        <TextField
                 id="checkout"
                  // label="checkout"
                  type="datetime-local"
                  value={checkout}
                  onChange={CheckouthandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                   shrink: true,
                 }}
                 style={{ width: 400 }}
                />&nbsp;&nbsp;&nbsp;&nbsp;
                {/* <TextField
                  id="timedateout"
                  //label="Alarm clock"
                  type="time"
                  value={timedateout}
                  onChange={TimedateouthandleChange}
                  className={classes.textField}
                   InputLabelProps={{
                  shrink: true,
                  }}
                     inputProps={{
                      step: 300, // 5 min
                   }}
                /> */}
                </div>
                </FormControl>
      </center>
                {/* <Typography align ="center"></Typography> */}
                <br></br>
                <div className={classes.paper}><strong>วันที่เวลารถเข้า</strong></div>
      <center>          <FormControl className={classes.margin} >
        <div>        <TextField
                 id="checkin"
                  // label="checkin"
                  type="datetime-local"
                  value={checkin}
                  onChange={CheckinhandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                   shrink: true,
                 }}
                 style={{ width: 400 }}
                />&nbsp;&nbsp;&nbsp;&nbsp;
                {/* <TextField
                  id="timedatein"
                  //label="Alarm clock"
                  type="time"
                  value={timedatein}
                  onChange={TimedateinhandleChange}
                  className={classes.textField}
                   InputLabelProps={{
                  shrink: true,
                  }}
                     inputProps={{
                      step: 300, // 5 min
                   }}
                /> */}
                </div>
                </FormControl>

                {/* <Typography align ="center"></Typography>    */}
        </center>
            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateCarcheckinout();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยัน
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/Carcheckinout"
                variant="contained"
              >
                กลับ
             </Button>
            </div>
          </form>
        </div>
             
      </Content>
    </Page>
  );
}
