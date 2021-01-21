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
  //เก็บข้อมูลที่จะดึงมา
  const [ambulances, setAmbulances] = useState<EntAmbulance[]>([]);
  const [purposes, setPurposes] = useState<EntPurpose[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);

  const [errormessege, setErrorMessege] = useState(String);
  const [alerttype, setAlertType] = useState(String);
  const [personError, setpersonError] = React.useState('');
  const [placeError, setplaceError] = React.useState('');
  const [distanceError, setdistanceError] = React.useState('');

  const [loading, setLoading] = useState(true);
  const [checkin, setCheckin] = useState(String);
  const [checkout, setCheckout] = useState(String);

  const [ambulanceid, setambulance] = useState(Number);
  const [purposeid, setpurpose] = useState(Number);
  const [userid, setUser] = useState(Number);
  const [note, setnote] = useState(String);
  const [place, setplace] = useState(String);
  const [person, setperson] = useState(Number);
  const [distance, setdistance] = useState(Number);

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

  const Placehandlehange = (event: any) => {
    const { value } = event.target;
    const validateValue = value
    checkfield('place', validateValue)
    setplace(event.target.value as string);
  };

  const Distancehandlehange = (event: any) => {
    const { value } = event.target;
    const validateValue = value
    checkfield('distance', validateValue)
    setdistance(event.target.value as number);
  };

  const Personhandlehange = (event: any) => {
    const { value } = event.target;
    const validateValue = value
    checkfield('person', validateValue)
    setperson(event.target.value as number);
  };

  const validatePerson = (val: number) =>{
    return val >=1 && val <=6 ? true:false
  }

  const validateDistance = (val: number) =>{
    return val >=10 ? true:false
  }

  const validatePlace = (val: string) =>{
    return val.length != 0 ? true:false
  }

  const checkfield = (id: string, value:any) =>{
    switch(id){
      case 'place':
        validatePlace(value) ? setplaceError('') : setplaceError('กรุณาใส่สถานที่');
        return;
      case 'distance':
        validateDistance(value) ? setdistanceError('') : setdistanceError('กรุณาใส่ระยะทาง(มากกว่า 10)ให้ถูกต้อง');
        return;
      case 'person':
        validatePerson(value) ? setpersonError('') : setpersonError('กรุณาใส่จำนวนเจ้าหน้าที่ให้ถูกต้อง(1-6คน/คัน)');
        return;
      default:
        return;
    }
  }

  const CreateCarcheckinout = async () => {
   if((checkin != "") && (checkout != "") && (checkin != null) && (checkout != null)) {
    const carcheckinouts = {
      ambulance: ambulanceid,
      name     : userid,
      purpose  : purposeid,
      person   : Number(person),
      distance : Number(distance),
      place    : place,
      note     : note,
      checkin  : checkin + ":00+07:00",
      checkout : checkout + ":00+07:00"
    };
    
    console.log(carcheckinouts);
    const apiUrl = 'http://localhost:8080/api/v1/carcheckinouts';
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(carcheckinouts),
        };
        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if (data.status === true) {
                    setErrorMessege("บันทึกข้อมูลสำเร็จ");
                    setAlertType("success");
                }
                else {
                    ErrorCaseCheck(data.error.Name);
                    setAlertType("error");
                }
            });}
        else{
          ErrorCaseCheck("กรุณาใส่เวลา");
          setAlertType("error");
        }
        setStatus(true);
    };

    const ErrorCaseCheck = (casename: string) => {
        if (casename == "person") { setErrorMessege("จำนวนเจ้าหน้าที่ 1-6 คน/คัน"); }
        else if (casename == "distance") { setErrorMessege("กรุณาใส่ระยะทางให้ถูกต้อง"); }
        else if (casename == "place") { setErrorMessege("กรุณาใส่สถานที่ก่อน"); }
        else { setErrorMessege("บันทึกไม่สำเร็จ"); }
    }

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
         
          <div>
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
           
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>จุดประสงค์</strong></div>
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
            <br></br><br></br>
            <div>
            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}      
              error = {personError ? true : false}
              helperText={personError}
              id="person"
              label="จำนวนเจ้าหน้าที่"
              variant="standard"
              color="secondary"
              type="number"
              size="medium"
              value={person}
              onChange={Personhandlehange}
            />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
             <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}    
              error = {distanceError ? true : false}
              helperText={distanceError}  
              id="distance"
              label="ระยะทาง"
              variant="standard"
              color="secondary"
              type="number"
              size="medium"
              value={distance}
              onChange={Distancehandlehange}
            />
            </div>
            <br></br><br></br>
            <div>
            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}   
              error = {placeError ? true : false}
              helperText={placeError}   
              id="place"
              label="สถานที่"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={place}
              onChange={Placehandlehange}
            />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}      
              id="note"
              label="หมายเหตุ"
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={note}
              onChange={Notehandlehange}
            />
            </div>
                <br></br><br></br>
            <div>
            <FormControl className={classes.margin} ><div className={classes.paper}><strong>วันที่เวลารถออก</strong></div>
                <TextField
                 id="checkout"
                 type="datetime-local"
                 value={checkout}
                 onChange={CheckouthandleChange}
                 className={classes.textField}
                 InputLabelProps={{
                 shrink: true,
                 }}
                 style={{ width: 400 }}
                />&nbsp;&nbsp;&nbsp;&nbsp;              
                </FormControl>
                <br></br>             
            <FormControl className={classes.margin} >  <div className={classes.paper}><strong>วันที่เวลารถเข้า</strong></div>
               <TextField
                 id="checkin"
                 type="datetime-local"
                 value={checkin}
                 onChange={CheckinhandleChange}
                 className={classes.textField}
                 InputLabelProps={{
                 shrink: true,
                 }}
                 style={{ width: 400 }}
                />&nbsp;&nbsp;&nbsp;&nbsp;  
                </FormControl>
            </div>
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
