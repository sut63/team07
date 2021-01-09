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
    EntCarInspection,
    EntCarInspectionFromJSON,
    EntCarInspectionFromJSONTyped,
    EntCarInspectionToJSON,
    EntJobPosition,
    EntJobPositionFromJSON,
    EntJobPositionFromJSONTyped,
    EntJobPositionToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntInspectionResultEdges
 */
export interface EntInspectionResultEdges {
    /**
     * Carinspections holds the value of the carinspections edge.
     * @type {Array<EntCarInspection>}
     * @memberof EntInspectionResultEdges
     */
    carinspections?: Array<EntCarInspection>;
    /**
     * 
     * @type {EntJobPosition}
     * @memberof EntInspectionResultEdges
     */
    jobposition?: EntJobPosition;
    /**
     * Statusof holds the value of the statusof edge.
     * @type {Array<EntAmbulance>}
     * @memberof EntInspectionResultEdges
     */
    statusof?: Array<EntAmbulance>;
}

export function EntInspectionResultEdgesFromJSON(json: any): EntInspectionResultEdges {
    return EntInspectionResultEdgesFromJSONTyped(json, false);
}

export function EntInspectionResultEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntInspectionResultEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'carinspections': !exists(json, 'Carinspections') ? undefined : ((json['Carinspections'] as Array<any>).map(EntCarInspectionFromJSON)),
        'jobposition': !exists(json, 'Jobposition') ? undefined : EntJobPositionFromJSON(json['Jobposition']),
        'statusof': !exists(json, 'Statusof') ? undefined : ((json['Statusof'] as Array<any>).map(EntAmbulanceFromJSON)),
    };
}

export function EntInspectionResultEdgesToJSON(value?: EntInspectionResultEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'carinspections': value.carinspections === undefined ? undefined : ((value.carinspections as Array<any>).map(EntCarInspectionToJSON)),
        'jobposition': EntJobPositionToJSON(value.jobposition),
        'statusof': value.statusof === undefined ? undefined : ((value.statusof as Array<any>).map(EntAmbulanceToJSON)),
    };
}

