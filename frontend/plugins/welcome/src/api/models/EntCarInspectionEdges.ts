/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntAmbulance,
    EntAmbulanceFromJSON,
    EntAmbulanceFromJSONTyped,
    EntAmbulanceToJSON,
    EntCarRepairrecord,
    EntCarRepairrecordFromJSON,
    EntCarRepairrecordFromJSONTyped,
    EntCarRepairrecordToJSON,
    EntInspectionResult,
    EntInspectionResultFromJSON,
    EntInspectionResultFromJSONTyped,
    EntInspectionResultToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCarInspectionEdges
 */
export interface EntCarInspectionEdges {
    /**
     * 
     * @type {EntAmbulance}
     * @memberof EntCarInspectionEdges
     */
    ambulance?: EntAmbulance;
    /**
     * Carrepairrecords holds the value of the carrepairrecords edge.
     * @type {Array<EntCarRepairrecord>}
     * @memberof EntCarInspectionEdges
     */
    carrepairrecords?: Array<EntCarRepairrecord>;
    /**
     * 
     * @type {EntInspectionResult}
     * @memberof EntCarInspectionEdges
     */
    inspectionresult?: EntInspectionResult;
    /**
     * 
     * @type {EntUser}
     * @memberof EntCarInspectionEdges
     */
    user?: EntUser;
}

export function EntCarInspectionEdgesFromJSON(json: any): EntCarInspectionEdges {
    return EntCarInspectionEdgesFromJSONTyped(json, false);
}

export function EntCarInspectionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCarInspectionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ambulance': !exists(json, 'Ambulance') ? undefined : EntAmbulanceFromJSON(json['Ambulance']),
        'carrepairrecords': !exists(json, 'Carrepairrecords') ? undefined : ((json['Carrepairrecords'] as Array<any>).map(EntCarRepairrecordFromJSON)),
        'inspectionresult': !exists(json, 'Inspectionresult') ? undefined : EntInspectionResultFromJSON(json['Inspectionresult']),
        'user': !exists(json, 'User') ? undefined : EntUserFromJSON(json['User']),
    };
}

export function EntCarInspectionEdgesToJSON(value?: EntCarInspectionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ambulance': EntAmbulanceToJSON(value.ambulance),
        'carrepairrecords': value.carrepairrecords === undefined ? undefined : ((value.carrepairrecords as Array<any>).map(EntCarRepairrecordToJSON)),
        'inspectionresult': EntInspectionResultToJSON(value.inspectionresult),
        'user': EntUserToJSON(value.user),
    };
}


