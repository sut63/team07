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
 * @interface EntJobPositionEdges
 */
export interface EntJobPositionEdges {
    /**
     * Inspectionresults holds the value of the inspectionresults edge.
     * @type {Array<EntInspectionResult>}
     * @memberof EntJobPositionEdges
     */
    inspectionresults?: Array<EntInspectionResult>;
    /**
     * Users holds the value of the users edge.
     * @type {Array<EntUser>}
     * @memberof EntJobPositionEdges
     */
    users?: Array<EntUser>;
}

export function EntJobPositionEdgesFromJSON(json: any): EntJobPositionEdges {
    return EntJobPositionEdgesFromJSONTyped(json, false);
}

export function EntJobPositionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntJobPositionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'inspectionresults': !exists(json, 'Inspectionresults') ? undefined : ((json['Inspectionresults'] as Array<any>).map(EntInspectionResultFromJSON)),
        'users': !exists(json, 'Users') ? undefined : ((json['Users'] as Array<any>).map(EntUserFromJSON)),
    };
}

export function EntJobPositionEdgesToJSON(value?: EntJobPositionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'inspectionresults': value.inspectionresults === undefined ? undefined : ((value.inspectionresults as Array<any>).map(EntInspectionResultToJSON)),
        'users': value.users === undefined ? undefined : ((value.users as Array<any>).map(EntUserToJSON)),
    };
}


