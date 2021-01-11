import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  ApiProvider,
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
import { EntCarInspection } from '../../api/models/EntCarInspection';
import { EntUser } from '../../api/models/EntUser';
import { EntRepairing } from '../../api/models/EntRepairing';
///import { EntCarInspectionEdges } from '../../api';
///import ComponentsTable from '../listambulance';

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

export default function CarInspectionPage() {
    const classes = useStyles();
    const profile = { givenName: 'ระบบบันทึกการบำรุงรักษา' };
    const api = new DefaultApi();

    const [status, setStatus] = useState(false);
    const [alert, setAlert] = useState(true);
    const [users, setUsers] = useState<EntUser[]>([]);
    const [carinspections, setCarinspections] = useState<EntCarInspection[]>([]);
    const [repairings, setRepairings] = useState<EntRepairing[]>([]);

    const [loading, setLoading] = useState(true);
    const [datetime, setDatetime] = useState(String);

    const [repairingid, setRepairing] = useState(Number);
    const [carinspectionid, setCarinspection] = useState(Number);
    const [userid, setUser] = useState(Number);

    useEffect(() => {
        const getCarinspections = async () => {
            const ci = await api.listCarinspection({ limit: 10, offset: 0 });
            setLoading(false);
            setCarinspections(ci);
        };
        getCarinspections();

        const getRepairings = async () => {
            const r = await api.listRepairing({limit:10, offset: 0});
            setLoading(false);
            setRepairings(r);
        };
        getRepairings();

        const getUsers = async () => {
            const us = await api.listUser();
            setLoading(false);
            setUsers(us);
        };
        getUsers();

        const checkJobPosition = async () => {
            const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
            setLoading(false);
            if (jobdata != "เจ้าหน้าที่ซ่อมบำรุงรถ" ) {
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
    const CarinspectionhandleChange = (event: React.ChangeEvent<{ value: unknown}>) => {
        setCarinspection(event.target.value as number)
    };

    const RepairinghandleChange = (event: React.ChangeEvent<{ value: unknown}>) => {
        setRepairing(event.target.value as number)
    };

    const UserhandleChange = (event: React.ChangeEvent<{ value: unknown}>) => {
        setUser(event.target.value as number)
    };

    const DateTimehandleChange = (event: any) => {
        setDatetime(event.target.value as string);
    };

    const CreateCarrepairrecord = async () => {
        if ((repairingid != null) && (carinspectionid != null) && (userid != null) && (datetime != "") && (datetime != null)) {
        const carrepairrecord = {
            carInspectionID : carinspectionid,
            repairingID : repairingid,
            userID : userid,
            datetime : datetime + ":00+07:00"
        };
        console.log(carrepairrecord);
        const res: any = await api.createCarrepairrecord({ carrepairrecord: carrepairrecord});
        setStatus(true);
        if(res.id != ''){
            setAlert(true);
            window.location.reload(false);
            }
        }else{
            setAlert(false);
            setStatus(true);
        }
        const timer = setTimeout(() => {
            setStatus(false);
        }, 1000);
    };

    return (
        <Page theme={pageTheme.home}>
            <Header title={`${profile.givenName}`}>
                
            </Header>
            <Content>
                <ContentHeader title="ลงข้อมูลซ่อมบำรุง">
                    {status ? (
                        <div>
                            {alert ? (
                                <Alert severity="success">
                                    บันทึกสำเร็จ
                                </Alert>
                            ) : (
                                <Alert severity="warning" style={{marginTop: 20}}>
                                    กรุณากรอกข้อมูลอีกครั้ง
                                </Alert>
                            )}
                        </div>
                    ) : null}
                </ContentHeader>
                <div className = {classes.root}>
                    <form noValidate autoComplete="off">
                        
                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <div className={classes.paper}><strong>รถพยาบาล</strong></div>
                                <InputLabel id="brand-lavel"></InputLabel>
                                <Select
                                    labelId="carinspection-label"
                                    id="carinspection"
                                    value={carinspectionid}
                                    onChange={CarinspectionhandleChange}
                                    style={{ width: 400}}
                                >
                                    {carinspections.filter((filter:any) => filter.edges?.inspectionresult?.resultName == "ส่งซ่อมแซม").map((item: any) => (
                                        <MenuItem value={item.id}>{item.edges?.ambulance?.carregistration}</MenuItem>
                                    ))}
                                </Select>
                            </FormControl>
                        </div>

                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <div className={classes.paper}><strong>ส่วนที่ซ่อม</strong></div>
                                <InputLabel id="repairing-label"></InputLabel>
                                <Select
                                    labelId="repairing-label"
                                    id="ส่วนที่ซ่อม"
                                    value={repairingid}
                                    onChange={RepairinghandleChange}
                                    style={{ width: 400}}
                                >
                                    {repairings.map((item: EntRepairing) => (
                                        <MenuItem value={item.id}>{item.repairpart}</MenuItem>
                                    ))}
                                </Select>
                            </FormControl>
                        </div>

                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <TextField
                                    id="user"
                                    label="เจ้าหน้าที่"
                                    type="string"
                                    size="medium"
                                    value={users.filter((filter:EntUser) => filter.id == userid).map((item:EntUser) => `${item.name} (${item.email})`)}
                                    style={{ width: 400 }}
                                />
                            </FormControl>
                        </div>

                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <TextField
                                    id="deathtime"
                                    label="เดือน/วัน/ปี"
                                    type="datetime-local"
                                    value={datetime}
                                    onChange={DateTimehandleChange}
                                    className={classes.textField}
                                    InputLabelProps={{
                                        shrink: true,
                                    }}
                                    style={{ width: 400 }}
                                />
                            </FormControl>
                        </div>

                        <center>
                        <div className={classes.margin}>
                            <Button
                                onClick={() => {
                                    CreateCarrepairrecord();
                                }}
                                variant="contained"
                                color="primary"
                            >
                                ยืนยัน
                            </Button>
                        </div>
                        </center>
                    </form>
                </div>
                
            </Content>
        </Page>   
    );
}