import React, { useState, useEffect } from 'react';
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
import Select from '@material-ui/core/Select';

import { EntUser } from '../../api/models/EntUser';
import { EntInspectionResult } from '../../api/models/EntInspectionResult';
import { EntAmbulance } from '../../api/models/EntAmbulance';

import CardMedia from '@material-ui/core/CardMedia';
import { Image1Base64Function } from '../../image/Image1';
import { Image2Base64Function } from '../../image/Image2';

import ComponanceTable from '../CarInspectionTable';

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


export default function CarInspectionPage() {
    const classes = useStyles();
    const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบตรวจสภาพรถ' };
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [alert, setAlert] = useState(true);
    const [users, setUsers] = useState<EntUser[]>([]);
    const [ambulances, setAmbulances] = useState<EntAmbulance[]>([]);
    const [inspectionresults, setInspectionResults] = useState<EntInspectionResult[]>([]);

    const [loading, setLoading] = useState(true);
    const [datetime, setDatetime] = useState(String);
    const [note, setNote] = useState(String);

    const [inspectionresultid, setInspectionResult] = useState(Number);
    const [ambulanceid, setAmbulance] = useState(Number);
    const [userid, setUser] = useState(Number);

    useEffect(() => {
        const getAmbulances = async () => {
            const res = await api.listAmbulance({ limit: 10, offset: 0 });
            setLoading(false);
            setAmbulances(res);
        };
        getAmbulances();

        const getInspectionResults = async () => {
            const res = await api.listInspectionresult();
            setLoading(false);
            setInspectionResults(res);
        };
        getInspectionResults();

        const getUsers = async () => {
            const res = await api.listUser();
            setLoading(false);
            setUsers(res);
        };
        getUsers();

        const checkJobPosition = async () => {
            const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
            setLoading(false);
            if (jobdata != "เจ้าหน้าที่ตรวจสภาพรถ" ) {
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


    const DateTimehandleChange = (event: any) => {
        setDatetime(event.target.value as string);
    };

    const NotehandleChange = (event: any) => {
        setNote(event.target.value as string);
    };

    const InspectionResulthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setInspectionResult(event.target.value as number);
    };

    const AmbulancehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setAmbulance(event.target.value as number);
    };

    const CreateCarInspection = async () => {
        const carinspection = {
            inspectionResultID: inspectionresultid,
            ambulanceID: ambulanceid,
            userID: userid,
            datetime: datetime + ":00+07:00",
            note: note


        };
        const res: any = await api.createCarinspection({ carinspection: carinspection });
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
        <Page theme={pageTheme.service}>
            <Header
                title={`${profile.givenName}`}
            //subtitle="Some quick intro and links."
            ></Header>
            <Content>
                <ContentHeader title="เพิ่มข้อมูลการตรวจสภาพรถ">
                    {status ? (
                        <div>
                            {alert ? (
                                <Alert severity="success">
                                    บันทึกสำเร็จ
                                </Alert>
                            ) : (
                                    <Alert severity="warning" style={{ marginTop: 20 }}>
                                        พบปัญหาระหว่างบันทึกข้อมูล
                                    </Alert>
                                )}
                        </div>
                    ) : null}
                </ContentHeader>
                <FormControl
                    className={classes.media}
                >
                    <CardMedia
                        component="img"
                        src={Image1Base64Function}
                    ></CardMedia>
                    <CardMedia
                        component="img"
                        src={Image2Base64Function}
                        style={ {marginTop:30} }
                    ></CardMedia>
                </FormControl>
                <div className={classes.root}>
                    <form noValidate autoComplete="off">

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

                        <FormControl
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel id="ambulance-label">รถพยาบาล</InputLabel>
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

                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <InputLabel id="result-label">ผลการตรวจสภาพ</InputLabel>
                                <Select
                                    labelId="result-label"
                                    id="result"
                                    value={inspectionresultid}
                                    onChange={InspectionResulthandleChange}
                                    style={{ width: 400 }}
                                >
                                    {inspectionresults.filter((filter: any) => filter.edges.jobposition.positionName == "เจ้าหน้าที่ตรวจสภาพรถ").map((item: EntInspectionResult) => (
                                        <MenuItem value={item.id}>{item.resultName}</MenuItem>
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
                                    id="note"
                                    label="หมายเหตุ"
                                    type="string"
                                    size="medium"
                                    value={note}
                                    onChange={NotehandleChange}
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
                                    id="datetime"
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

                        <div className={classes.margin}>
                            <Button
                                onClick={() => {
                                    CreateCarInspection();
                                }}
                                variant="contained"
                                color="primary"
                            >
                                ยืนยัน
                            </Button>
                        </div>
                    </form>
                </div>
                <ComponanceTable></ComponanceTable>
            </Content>
        </Page>
    );
}