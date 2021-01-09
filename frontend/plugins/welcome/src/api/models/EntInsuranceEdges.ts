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
} from './';

/**
 * 
 * @export
 * @interface EntInsuranceEdges
 */
export interface EntInsuranceEdges {
    /**
     * Insuranceof holds the value of the insuranceof edge.
     * @type {Array<EntAmbulance>}
     * @memberof EntInsuranceEdges
     */
    insuranceof?: Array<EntAmbulance>;
}

export function EntInsuranceEdgesFromJSON(json: any): EntInsuranceEdges {
    return EntInsuranceEdgesFromJSONTyped(json, false);
}

export function EntInsuranceEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntInsuranceEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'insuranceof': !exists(json, 'Insuranceof') ? undefined : ((json['Insuranceof'] as Array<any>).map(EntAmbulanceFromJSON)),
    };
}

export function EntInsuranceEdgesToJSON(value?: EntInsuranceEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'insuranceof': value.insuranceof === undefined ? undefined : ((value.insuranceof as Array<any>).map(EntAmbulanceToJSON)),
    };
}

