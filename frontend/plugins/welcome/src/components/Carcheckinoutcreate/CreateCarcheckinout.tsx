import React, { useState, useEffect } from 'react';
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

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
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

  const UserthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };
  const Notehandlehange = (event: any) => {
    setnote(event.target.value as string);
  };

  const CreateCarcheckinout = async () => {
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
      window.location.reload(false);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };
  

  return (
 <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      //subtitle="Some quick intro and links."
      ></Header>
      <Content>
        <ContentHeader title="ลงทะเบียนรถเข้าออก">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    กรุณากรอกข้อมูลอีกครั้ง
                  </Alert>
                )}
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
            </div>
            
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>เจ้าหน้าที่รถพยาบาล</strong></div>
                <InputLabel id="user-label"></InputLabel>
                <Select
                  labelId="user-label"
                  id="user"
                  value={userid}
                  onChange={UserthandleChange}
                  style={{ width: 400 }}
                >
                  {users.map((item: EntUser) => (
                  <MenuItem value={item.id}>{item.name}</MenuItem>
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
          
            <div className={classes.paper}><strong>วันที่เวลารถออก</strong></div>
                <TextField
                    className={classes.formControl}
                    id="checkout"
                    type="datetime-local"
                    value={checkout}
                    onChange={CheckouthandleChange}
                    InputLabelProps={{
                     shrink: true,
                     }}
                 />
                <Typography align ="right"></Typography>

                <div className={classes.paper}><strong>วันที่เวลารถเข้า</strong></div>
                <TextField
                    className={classes.formControl}
                    id="checkin"
                    type="datetime-local"
                    value={checkin}
                    onChange={CheckinhandleChange}
                    InputLabelProps={{
                     shrink: true,
                     }}
                 />
                <Typography align ="right"></Typography>   

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
