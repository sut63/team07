import React, { useState, useEffect } from 'react';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';

import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { Alert } from '@material-ui/lab';

import moment from 'moment';

import { EntInspectionResult } from '../../api/models/EntInspectionResult';

import SearchIcon from '@material-ui/icons/Search';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        table: {
            minWidth: 650,
        },
        margin: {
            margin: theme.spacing(3),
        },
    }),
);

export default function CarInspectionSearchPage() {
    const classes = useStyles();
    const api = new DefaultApi();
    const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบตรวจสภาพรถ' };
    const [carinspections, setCarInspections] = useState([]);
    const [loading, setLoading] = useState(true);
    const [status, setStatus] = useState(false);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [inspectionresults, setInspectionResults] = useState<EntInspectionResult[]>([]);

    const [inspectionresultsearch, setInspectionResult] = useState(Number);
    const [usersearch, setUserSearch] = useState(String);
    const [ambulancesearch, setAmbulanceSearch] = useState(String);

    useEffect(() => {
        const getInspectionResults = async () => {
            const res = await api.listInspectionresult();
            setLoading(false);
            setInspectionResults(res);
        };
        getInspectionResults();

        const checkJobPosition = async () => {
            const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
            setLoading(false);
            if (jobdata != "เจ้าหน้าที่ตรวจสภาพรถ") {
                localStorage.setItem("userdata", JSON.stringify(null));
                localStorage.setItem("jobpositiondata", JSON.stringify(null));
                history.pushState("", "", "./");
                window.location.reload(false);
            }
        }
        checkJobPosition();


    }, [loading]);

    const SearchCarInspection = async () => {
        const apiUrl = `http://localhost:8080/api/v1/searchcarinspections?ambulance=${ambulancesearch}&result=${inspectionresultsearch}&user=${usersearch}`;
        const requestOptions = {
            method: 'GET',
        };
        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data.data)
                setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
                setAlertType("error");
                setCarInspections([]);
                if (data.data != null) {
                    if(data.data.length >= 1) {
                        setErrorMessege("พบข้อมูลที่ค้นหา");
                        setAlertType("success");
                        console.log(data.data)
                        setCarInspections(data.data);
                    }
                }

                setStatus(true);
            });

    }

    const InspectionResulthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setInspectionResult(event.target.value as number);
    };

    const UserSearchhandleChange = (event: any) => {
        setUserSearch(event.target.value as string);
    };

    const AmbulanceSearchhandleChange = (event: any) => {
        setAmbulanceSearch(event.target.value as string);
    };

    return (
        <Page theme={pageTheme.service}>
            <Header
                title={`${profile.givenName}`}
            //subtitle="Some quick intro and links."
            ></Header>
            <Content>
                <ContentHeader title="ค้นหาข้อมูลการตรวจสภาพรถ">
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
                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <TextField
                        id="search"
                        label="ค้นหาผู้บันทึกด้วย keyword"
                        type="string"
                        size="medium"
                        value={usersearch}
                        onChange={UserSearchhandleChange}
                        style={{ width: 200 }}
                    />
                </FormControl>

                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <TextField
                        id="search"
                        label="ค้นหาเลขทะเบียนรถด้วย keyword"
                        type="string"
                        size="medium"
                        value={ambulancesearch}
                        onChange={AmbulanceSearchhandleChange}
                        style={{ width: 300 }}
                    />
                </FormControl>

                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <InputLabel id="result-label">ค้นหาผลการตรวจสภาพรถ</InputLabel>
                    <Select
                        labelId="result-label"
                        id="result"
                        value={inspectionresultsearch}
                        onChange={InspectionResulthandleChange}
                        style={{ width: 200 }}

                    >   <MenuItem value={0}>-</MenuItem>
                        {inspectionresults.filter((filter: EntInspectionResult) => filter.edges?.jobposition?.positionName == "เจ้าหน้าที่ตรวจสภาพรถ").map((item: EntInspectionResult) => (
                            <MenuItem value={item.id}>{item.resultName}</MenuItem>
                        ))}
                    </Select>
                </FormControl>

                <div className={classes.margin}>
                    <Button
                        onClick={() => {
                            SearchCarInspection();
                        }}
                        variant="contained"
                        color="primary"
                        style={{ width: 180, backgroundColor: "#365391" }}
                        startIcon={<SearchIcon />}
                    >
                        ค้นหา
                    </Button>
                    <Button
                        onClick={() => {
                            history.pushState("", "", "./carinspection");
                            window.location.reload(false);
                        }}
                        variant="contained"
                        color="primary"
                        style={{ width: 180, marginLeft: 40, backgroundColor: "#a31f1f" }}
                        startIcon={<ArrowBackIcon />}
                    >
                        กลับ
                    </Button>
                </div>

                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">เจ้าหน้าที่</TableCell>
                                <TableCell align="center">รถพยาบาล</TableCell>
                                <TableCell align="center">ศูนย์ล้อ (เมตร)</TableCell>
                                <TableCell align="center">ระดับเสียง (เดซิเบล)</TableCell>
                                <TableCell align="center">ควันดำ (เปอร์เซ็นต์)</TableCell>
                                <TableCell align="center">ผลการตรวจสภาพรถ</TableCell>
                                <TableCell align="center">วัน/เดือน/ปี เวลา</TableCell>
                                <TableCell align="center">หมายเหตุ</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {carinspections.map((item: any) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges.User.name}</TableCell>
                                    <TableCell align="center">{item.edges.Ambulance.carregistration}</TableCell>
                                    <TableCell align="center">{item.wheel_center}</TableCell>
                                    <TableCell align="center">{item.sound_level}</TableCell>
                                    <TableCell align="center">{item.blacksmoke}</TableCell>
                                    <TableCell align="center">{item.edges.Inspectionresult.result_name}</TableCell>
                                    <TableCell align="center">{moment(item.datetime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                                    <TableCell align="center">{item.note}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Content>
        </Page>
    );
}