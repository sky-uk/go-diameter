// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file is auto-generated from our dictionaries.

package dict

import (
	"bytes"
	"fmt"
)

// Default is a Parser object with pre-loaded
// Base Protocol and Credit Control dictionaries.
var Default *Parser

func init() {
	var dictionaries = []struct{ name, xml string }{
		{"Base", baseXML},
		{"Credit Control", creditcontrolXML},
		{"Gx Charging Control", gxcreditcontrolXML},
		/*{"Network Access Server", networkaccessserverXML},*/
		/*{"TGPP", tgpprorfXML},*/
		/*{"TGPP_S6a", tgpps6aXML},*/
		/*{"TGPP_Swx", tgppswxXML},*/
		{"TGPP_CLF", tgppe2XML},
	}
	var err error
	Default, err = NewParser()
	if err != nil {
		panic(err)
	}
	for _, dict := range dictionaries {
		err = Default.Load(bytes.NewReader([]byte(dict.xml)))
		if err != nil {
			panic(fmt.Sprintf("Cannot load %s dictionary: %s", dict.name, err))
		}
	}
}

var baseXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

	<application id="0" name="Base"> <!-- Diameter Common Messages -->

		<command code="257" short="CE" name="Capabilities-Exchange">
			<request>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Host-IP-Address" required="true" min="1"/>
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="Product-Name" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Supported-Vendor-Id" required="False"/>
				<rule avp="Auth-Application-Id" required="False"/>
				<rule avp="Inband-Security-Id" required="False"/>
				<rule avp="Acct-Application-Id" required="False"/>
				<rule avp="Vendor-Specific-Application-Id" required="False"/>
				<rule avp="Firmware-Revision" required="False" max="1"/>
			</request>
			<answer>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Host-IP-Address" required="true" min="1"/>
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="Product-Name" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Supported-Vendor-Id" required="False"/>
				<rule avp="Auth-Application-Id" required="False"/>
				<rule avp="Inband-Security-Id" required="False"/>
				<rule avp="Acct-Application-Id" required="False"/>
				<rule avp="Vendor-Specific-Application-Id" required="False"/>
				<rule avp="Firmware-Revision" required="False" max="1"/>
			</answer>
		</command>

		<command code="258" short="RA" name="Re-Auth">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Re-Auth-Request-Type" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="271" short="AC" name="Accounting">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Accounting-Record-Type" required="true" max="1"/>
				<rule avp="Accounting-Record-Number" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="false" max="1"/>
				<rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Accounting-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Interim-Interval" required="false" max="1"/>
				<rule avp="Accounting-Realtime-Required" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Accounting-Record-Type" required="true" max="1"/>
				<rule avp="Accounting-Record-Number" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="false" max="1"/>
				<rule avp="Vendor-Specific-Application-Id" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Accounting-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Acct-Interim-Interval" required="false" max="1"/>
				<rule avp="Accounting-Realtime-Required" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="274" short="AS" name="Abort-Session">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="275" short="ST" name="Session-Termination">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Termination-Cause" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
				<rule avp="Route-Record" required="false"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="Class" required="false"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Error-Reporting-Host" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false"/>
			</answer>
		</command>

		<command code="280" short="DW" name="Device-Watchdog">
			<request>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
			</request>
			<answer>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
			</answer>
		</command>

		<command code="282" short="DP" name="Disconnect-Peer">
			<request>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Disconnect-Cause" required="false" max="1"/>
			</request>
			<answer>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
			</answer>
		</command>

		<avp name="Acct-Interim-Interval" code="85" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Accounting-Realtime-Required" code="483" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Enumerated">
				<item code="1" name="DELIVER_AND_GRANT"/>
				<item code="2" name="GRANT_AND_STORE"/>
				<item code="3" name="GRANT_AND_LOSE"/>
			</data>
		</avp>

		<avp name="Acct-Multi-Session-Id" code="50" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Accounting-Record-Number" code="485" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Accounting-Record-Type" code="480" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Enumerated">
				<item code="1" name="EVENT_RECORD"/>
				<item code="2" name="START_RECORD"/>
				<item code="3" name="INTERIM_RECORD"/>
				<item code="4" name="STOP_RECORD"/>
			</data>
		</avp>

		<avp name="Accounting-Session-Id" code="44" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="OctetString"/>
		</avp>

		<avp name="Accounting-Sub-Session-Id" code="287" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned64"/>
		</avp>

		<avp name="Acct-Application-Id" code="259" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Application-Id" code="258" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Request-Type" code="274" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="1" name="AUTHENTICATE_ONLY"/>
				<item code="2" name="AUTHORIZE_ONLY"/>
				<item code="3" name="AUTHORIZE_AUTHENTICATE"/>
			</data>
		</avp>

		<avp name="Authorization-Lifetime" code="291" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Grace-Period" code="276" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Auth-Session-State" code="277" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="STATE_MAINTAINED"/>
				<item code="1" name="NO_STATE_MAINTAINED"/>
			</data>
		</avp>

		<avp name="Re-Auth-Request-Type" code="285" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="AUTHORIZE_ONLY"/>
				<item code="1" name="AUTHORIZE_AUTHENTICATE"/>
			</data>
		</avp>

		<avp name="Class" code="25" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="OctetString"/>
		</avp>

		<avp name="Destination-Host" code="293" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Destination-Realm" code="283" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Disconnect-Cause" code="273" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="REBOOTING"/>
				<item code="1" name="BUSY"/>
				<item code="2" name="DO_NOT_WANT_TO_TALK_TO_YOU"/>
			</data>
		</avp>

		<avp name="Error-Message" code="281" must="-" may="P" must-not="V,M" may-encrypt="-">
			<data type="UTF8String"/>
		</avp>

		<avp name="Error-Reporting-Host" code="294" must="-" may="P" must-not="V,M" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Event-Timestamp" code="55" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Time"/>
		</avp>

		<avp name="Experimental-Result" code="297" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Vendor-Id" required="true" max="1"/>
				<rule avp="Experimental-Result-Code" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Experimental-Result-Code" code="298" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Failed-AVP" code="279" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Grouped"/>
		</avp>

		<avp name="Firmware-Revision" code="267" must="-" may="-" must-not="P,V,M" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Host-IP-Address" code="257" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Address"/>
		</avp>

		<avp name="Inband-Security-Id" code="299" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Multi-Round-Time-Out" code="272" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Origin-Host" code="264" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Origin-Realm" code="296" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Origin-State-Id" code="278" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Product-Name" code="269" must="-" may="-" must-not="P,V,M" may-encrypt="-">
			<data type="UTF8String"/>
		</avp>

		<avp name="Proxy-Host" code="280" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Proxy-Info" code="284" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Proxy-Host" required="true" max="1"/>
				<rule avp="Proxy-State" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Proxy-State" code="33" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="OctetString"/>
		</avp>

		<avp name="Redirect-Host" code="292" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="DiameterURI"/>
		</avp>

		<avp name="Redirect-Host-Usage" code="261" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="DONT_CACHE"/>
				<item code="1" name="ALL_SESSION"/>
				<item code="2" name="ALL_REALM"/>
				<item code="3" name="REALM_AND_APPLICATION"/>
				<item code="4" name="ALL_APPLICATION"/>
				<item code="5" name="ALL_HOST"/>
				<item code="6" name="ALL_USER"/>
			</data>
		</avp>

		<avp name="Redirect-Max-Cache-Time" code="262" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Result-Code" code="268" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Route-Record" code="282" must="M" may="-" must-not="P,V" may-encrypt="-">
			<data type="DiameterIdentity"/>
		</avp>

		<avp name="Session-Id" code="263" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Session-Timeout" code="27" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Session-Binding" code="270" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Session-Server-Failover" code="271" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Enumerated">
				<item code="0" name="REFUSE_SERVICE"/>
				<item code="1" name="TRY_AGAIN"/>
				<item code="2" name="ALLOW_SERVICE"/>
				<item code="3" name="TRY_AGAIN_ALLOW_SERVICE"/>
			</data>
		</avp>

		<avp name="Supported-Vendor-Id" code="265" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Termination-Cause" code="295" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="1" name="DIAMETER_LOGOUT"/>
				<item code="2" name="DIAMETER_SERVICE_NOT_PROVIDED"/>
				<item code="3" name="DIAMETER_BAD_ANSWER"/>
				<item code="4" name="DIAMETER_ADMINISTRATIVE"/>
				<item code="5" name="DIAMETER_LINK_BROKEN"/>
				<item code="6" name="DIAMETER_AUTH_EXPIRED"/>
				<item code="7" name="DIAMETER_USER_MOVED"/>
				<item code="8" name="DIAMETER_SESSION_TIMEOUT"/>
			</data>
		</avp>

		<avp name="User-Name" code="1" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="UTF8String"/>
		</avp>

		<avp name="Vendor-Id" code="266" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Unsigned32"/>
		</avp>

		<avp name="Vendor-Specific-Application-Id" code="260" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Grouped">
				<rule avp="Vendor-Id" required="false" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Acct-Application-Id" required="true" max="1"/>
			</data>
		</avp>

		<!-- IETF RFC 7683 - https://tools.ietf.org/html/rfc7683 -->
		<avp name="OC-Supported-Features" code="621" must-not="V">
			<data type="Grouped">
				<rule avp="OC-Feature-Vector" required="false"/>
				<rule avp="AVP" required="false"/>
			</data>
		</avp>

		<avp name="OC-Feature-Vector" code="622" must-not="V">
			<data type="Unsigned64"/>
		</avp>

		<avp name="OC-OLR" code="623" must-not="V">
			<data type="Grouped">
				<rule avp="OC-Sequence-Number" required="true" max="1"/>
				<rule avp="OC-Report-Type" required="true" max="1"/>
				<rule avp="OC-Reduction-Percentage" required="false" max="1"/>
				<rule avp="OC-Validity-Duration" required="false" max="1"/>
				<rule avp="AVP" required="false"/>
			</data>
		</avp>

		<avp name="OC-Sequence-Number" code="624" must-not="V">
			<data type="Unsigned64"/>
		</avp>

		<avp name="OC-Validity-Duration" code="625" must-not="V">
			<data type="Unsigned32"/>
		</avp>

		<avp name="OC-Report-Type" code="626" must-not="V">
			<data type="Enumerated">
				<item code="0" name="HOST_REPORT"/>
				<item code="1" name="REALM_REPORT"/>
			</data>
		</avp>

		<avp name="OC-Reduction-Percentage" code="627" must-not="V">
			<data type="Unsigned32"/>
		</avp>

		<!-- IETF RFC 7944 - https://tools.ietf.org/html/rfc7944 -->
		<avp name="DRMP" code="301" must-not="V">
			<data type="Enumerated">
				<item code="0" name="PRIORITY_0"/>
				<item code="1" name="PRIORITY_1"/>
				<item code="2" name="PRIORITY_2"/>
				<item code="3" name="PRIORITY_3"/>
				<item code="4" name="PRIORITY_4"/>
				<item code="5" name="PRIORITY_5"/>
				<item code="6" name="PRIORITY_6"/>
				<item code="7" name="PRIORITY_7"/>
				<item code="8" name="PRIORITY_8"/>
				<item code="9" name="PRIORITY_9"/>
				<item code="10" name="PRIORITY_10"/>
				<item code="11" name="PRIORITY_11"/>
				<item code="12" name="PRIORITY_12"/>
				<item code="13" name="PRIORITY_13"/>
				<item code="14" name="PRIORITY_14"/>
				<item code="15" name="PRIORITY_15"/>
			</data>
		</avp>


	</application>
	<application id="3" type="acct" name="Base Accounting"> <!-- Diameter Base Accounting Messages -->
	</application>
</diameter>`

var creditcontrolXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

	<application id="4" type="auth" name="Charging Control">
		<!-- Diameter Credit Control Application -->
		<!-- http://tools.ietf.org/html/rfc4006 -->

		<command code="272" short="CC" name="Credit-Control">
			<request>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.1 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Service-Context-Id" required="true" max="1"/>
				<rule avp="CC-Request-Type" required="true" max="1"/>
				<rule avp="CC-Request-Number" required="true" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="CC-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Subscription-Id" required="false" max="1"/>
				<rule avp="Service-Identifier" required="false" max="1"/>
				<rule avp="Termination-Cause" required="false" max="1"/>
				<rule avp="Requested-Service-Unit" required="false" max="1"/>
				<rule avp="Requested-Action" required="false" max="1"/>
				<rule avp="Used-Service-Unit" required="false" max="1"/>
				<rule avp="Multiple-Services-Indicator" required="false" max="1"/>
				<rule avp="Multiple-Services-Credit-Control" required="false" max="1"/>
				<rule avp="Service-Parameter-Info" required="false" max="1"/>
				<rule avp="CC-Correlation-Id" required="false" max="1"/>
				<rule avp="User-Equipment-Info" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
				<rule avp="Service-Information" required="false" max="1"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.2 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="CC-Request-Type" required="true" max="1"/>
				<rule avp="CC-Request-Number" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
				<rule avp="CC-Session-Failover" required="false" max="1"/>
				<rule avp="CC-Sub-Session-Id" required="false" max="1"/>
				<rule avp="Acct-Multi-Session-Id" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Event-Timestamp" required="false" max="1"/>
				<rule avp="Granted-Service-Unit" required="false" max="1"/>
				<rule avp="Multiple-Services-Credit-Control" required="false" max="1"/>
				<rule avp="Cost-Information" required="false" max="1"/>
				<rule avp="Final-Unit-Indication" required="false" max="1"/>
				<rule avp="Check-Balance-Result" required="false" max="1"/>
				<rule avp="Credit-Control-Failure-Handling" required="false" max="1"/>
				<rule avp="Direct-Debiting-Failure-Handling" required="false" max="1"/>
				<rule avp="Validity-Time" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false" max="1"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
			</answer>
		</command>

		<avp name="CC-Correlation-Id" code="411" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.1 -->
			<data type="OctetString"/>
		</avp>

		<avp name="CC-Input-Octets" code="412" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.24 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Money" code="413" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.22 -->
			<data type="Grouped">
				<rule avp="Unit-Value" required="true" max="1"/>
				<rule avp="Currency-Code" required="true" max="1"/>
			</data>
		</avp>

		<avp name="CC-Output-Octets" code="414" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.25 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Request-Number" code="415" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.2 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="CC-Request-Type" code="416" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.3 -->
			<data type="Enumerated">
				<item code="1" name="INITIAL_REQUEST"/>
				<item code="2" name="UPDATE_REQUEST"/>
				<item code="3" name="TERMINATION_REQUEST"/>
			</data>
		</avp>

		<avp name="CC-Service-Specific-Units" code="417" must="M" may="P" must-not="V" may-encrypt="Y">
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Session-Failover" code="418" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.4 -->
			<data type="Enumerated">
				<item code="0" name="FAILOVER_NOT_SUPPORTED"/>
				<item code="1" name="FAILOVER_SUPPORTED"/>
			</data>
		</avp>

		<avp name="CC-Sub-Session-Id" code="419" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.5 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Time" code="420" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.21 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="CC-Total-Octets" code="421" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.23 -->
			<data type="Unsigned64"/>
		</avp>

		<avp name="CC-Unit-Type" code="454" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.32 -->
			<data type="Enumerated">
				<item code="0" name="TIME"/>
				<item code="1" name="MONEY"/>
				<item code="2" name="TOTAL-OCTETS"/>
				<item code="3" name="INPUT-OCTETS"/>
				<item code="4" name="OUTPUT-OCTETS"/>
				<item code="5" name="SERVICE-SPECIFIC-UNITS"/>
			</data>
		</avp>

		<avp name="Check-Balance-Result" code="422" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.6 -->
			<data type="Enumerated">
				<item code="0" name="ENOUGH_CREDIT"/>
				<item code="1" name="NO_CREDIT"/>
			</data>
		</avp>

		<avp name="Cost-Information" code="423" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.7 -->
			<data type="Grouped">
				<rule avp="Unit-Value" required="true" max="1"/>
				<rule avp="Currency-Code" required="true" max="1"/>
				<rule avp="Cost-Unit" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Cost-Unit" code="424" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.12 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Credit-Control" code="426" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.13 -->
			<data type="Enumerated">
				<item code="0" name="CREDIT_AUTHORIZATION"/>
				<item code="1" name="RE_AUTHORIZATION"/>
			</data>
		</avp>

		<avp name="Credit-Control-Failure-Handling" code="427" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.14 -->
			<data type="Enumerated">
				<item code="0" name="TERMINATE"/>
				<item code="1" name="CONTINUE"/>
				<item code="2" name="RETRY_AND_TERMINATE"/>
			</data>
		</avp>

		<avp name="Currency-Code" code="425" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.11 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Direct-Debiting-Failure-Handling" code="428" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.15 -->
			<data type="Enumerated">
				<item code="0" name="TERMINATE_OR_BUFFER"/>
				<item code="1" name="CONTINUE"/>
			</data>
		</avp>

		<avp name="Exponent" code="429" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.9 -->
			<data type="Integer32"/>
		</avp>

		<avp name="Final-Unit-Action" code="449" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.35 -->
			<data type="Enumerated">
				<item code="0" name="TERMINATE"/>
				<item code="1" name="REDIRECT"/>
				<item code="2" name="RESTRICT_ACCESS"/>
			</data>
		</avp>

		<avp name="Final-Unit-Indication" code="430" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.34 -->
			<data type="Grouped">
				<rule avp="Final-Unit-Action" required="true" max="1"/>
				<rule avp="Restriction-Filter-Rule" required="false" max="1"/>
				<rule avp="Filter-Id" required="false" max="1"/>
				<rule avp="Redirect-Server" required="false" max="1"/>
			</data>
		</avp>

		<avp name="Granted-Service-Unit" code="431" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.17 -->
			<data type="Grouped">
				<rule avp="Tariff-Time-Change" required="false" max="1"/>
				<rule avp="CC-Time" required="false" max="1"/>
				<rule avp="CC-Money" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="G-S-U-Pool-Identifier" code="453" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.31 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="G-S-U-Pool-Reference" code="457" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.30 -->
			<data type="Grouped">
				<rule avp="G-S-U-Pool-Identifier" required="true" max="1"/>
				<rule avp="CC-Unit-Type" required="true" max="1"/>
				<rule avp="Unit-Value" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Multiple-Services-Credit-Control" code="456" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.16 -->
			<data type="Grouped">
				<rule avp="Granted-Service-Unit" required="false" max="1"/>
				<rule avp="Requested-Service-Unit" required="false" max="1"/>
				<rule avp="Used-Service-Unit" required="false" max="1"/>
				<rule avp="Tariff-Change-Usage" required="false" max="1"/>
				<rule avp="Service-Identifier" required="false" max="1"/>
				<rule avp="Rating-Group" required="false" max="1"/>
				<rule avp="G-S-U-Pool-Reference" required="false" max="1"/>
				<rule avp="Validity-Time" required="false" max="1"/>
				<rule avp="Result-Code" required="false" max="1"/>
				<rule avp="Final-Unit-Indication" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="Multiple-Services-Indicator" code="455" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.40 -->
			<data type="Enumerated">
				<item code="0" name="MULTIPLE_SERVICES_NOT_SUPPORTED"/>
				<item code="1" name="MULTIPLE_SERVICES_SUPPORTED"/>
			</data>
		</avp>

		<avp name="Rating-Group" code="432" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.29 -->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Redirect-Address-Type" code="433" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.38 -->
			<data type="Enumerated">
				<item code="0" name="IPv4 Address"/>
				<item code="1" name="IPv6 Address"/>
				<item code="2" name="URL"/>
				<item code="3" name="SIP URI"/>
			</data>
		</avp>

		<avp name="Redirect-Server" code="434" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.37 -->
			<data type="Grouped">
				<rule avp="Redirect-Address-Type" required="true" max="1"/>
				<rule avp="Redirect-Server-Address" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Redirect-Server-Address" code="435" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.39 -->
			<data type="UTF8String"/>
		</avp>

		<avp name="Requested-Action" code="436" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.41 -->
			<data type="Enumerated">
				<item code="0" name="DIRECT_DEBITING"/>
				<item code="1" name="REFUND_ACCOUNT"/>
				<item code="2" name="CHECK_BALANCE"/>
				<item code="3" name="PRICE_ENQUIRY"/>
			</data>
		</avp>

		<avp name="Requested-Service-Unit" code="437" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.18-->
			<data type="Grouped">
				<rule avp="CC-Time" required="false" max="1"/>
				<rule avp="CC-Money" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="Restriction-Filter-Rule" code="438" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.36-->
			<data type="IPFilterRule"/>
		</avp>

		<avp name="Service-Context-Id" code="461" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.42-->
			<data type="UTF8String"/>
		</avp>

		<avp name="Service-Identifier" code="439" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.28-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Service-Parameter-Info" code="440" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.43-->
			<data type="Grouped">
				<rule avp="Service-Parameter-Type" required="true" max="1"/>
				<rule avp="Service-Parameter-Value" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Service-Parameter-Type" code="441" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.44-->
			<data type="Unsigned32"/>
		</avp>

		<avp name="Service-Parameter-Value" code="442" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.45-->
			<data type="OctetString"/>
		</avp>

		<avp name="Subscription-Id" code="443" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.46-->
			<data type="Grouped">
				<rule avp="Subscription-Id-Type" required="true" max="1"/>
				<rule avp="Subscription-Id-Data" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Subscription-Id-Data" code="444" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.48-->
			<data type="UTF8String"/>
		</avp>

		<avp name="Subscription-Id-Type" code="450" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.47-->
			<data type="Enumerated">
				<item code="0" name="END_USER_E164"/>
				<item code="1" name="END_USER_IMSI"/>
				<item code="2" name="END_USER_SIP_URI"/>
				<item code="3" name="END_USER_NAI"/>
			</data>
		</avp>

		<avp name="Tariff-Change-Usage" code="452" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.27-->
			<data type="Enumerated">
				<item code="0" name="UNIT_BEFORE_TARIFF_CHANGE"/>
				<item code="1" name="UNIT_AFTER_TARIFF_CHANGE"/>
				<item code="2" name="UNIT_INDETERMINATE"/>
			</data>
		</avp>

		<avp name="Tariff-Time-Change" code="451" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.20-->
			<data type="Time"/>
		</avp>

		<avp name="Unit-Value" code="445" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.8-->
			<data type="Grouped">
				<rule avp="Value-Digits" required="true" max="1"/>
				<rule avp="Exponent" required="true" max="1"/>
			</data>
		</avp>

		<avp name="Used-Service-Unit" code="446" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.19-->
			<data type="Grouped">
				<rule avp="Tariff-Change-Usage" required="false" max="1"/>
				<rule avp="CC-Time" required="false" max="1"/>
				<rule avp="CC-Money" required="false" max="1"/>
				<rule avp="CC-Total-Octets" required="false" max="1"/>
				<rule avp="CC-Input-Octets" required="false" max="1"/>
				<rule avp="CC-Output-Octets" required="false" max="1"/>
				<rule avp="CC-Service-Specific-Units" required="false" max="1"/>
				<!-- *[ AVP ]-->
			</data>
		</avp>

		<avp name="User-Equipment-Info" code="458" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.49-->
			<data type="Grouped">
				<rule avp="User-Equipment-Info-Type" required="true" max="1"/>
				<rule avp="User-Equipment-Info-Value" required="true" max="1"/>
			</data>
		</avp>

		<avp name="User-Equipment-Info-Type" code="459" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.50-->
			<data type="Enumerated">
				<item code="0" name="IMEISV"/>
				<item code="1" name="MAC"/>
				<item code="2" name="EUI64"/>
				<item code="3" name="MODIFIED_EUI64"/>
			</data>
		</avp>

		<avp name="User-Equipment-Info-Value" code="460" must="-" may="P,M" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.51-->
			<data type="OctetString"/>
		</avp>

		<avp name="Value-Digits" code="447" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.10-->
			<data type="Integer64"/>
		</avp>

		<avp name="Validity-Time" code="448" must="M" may="P" must-not="V" may-encrypt="Y">
			<!-- http://tools.ietf.org/html/rfc4006#section-8.33-->
			<data type="Unsigned32"/>
		</avp>
	</application>
</diameter>`

var diametersyXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

	<application id="16777302" type="auth" name="Diameter Sy">
		<!-- Diameter Credit Control Application -->
		<!-- http://tools.ietf.org/html/rfc4006 -->

		<command code="8388635" short="SL" name="Spending-Limit">
			<request>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.1 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="SL-Request-Type" required="true" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Subscription-Id" required="false" max="1"/>
				<rule avp="Policy-Counter-Identifier" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
				<rule avp="Service-Information" required="false" max="1"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.2 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false" max="1"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
			</answer>
		</command>

		<avp name="SL-Request-Type" code="2904" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="INITIAL_REQUEST"/>
				<item code="1" name="INTERMEDIATE_REQUEST"/>
			</data>
		</avp>
		
    </application>
</diameter>`

var gxcreditcontrolXML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

    <application id="16777238" type="auth" name="Gx Charging Control">
        <!-- Diameter Gx Credit Control Application -->
        <!-- 3GPP 29.212 -->

        <vendor id="10415" name="TGPP"/>
        <command code="272" short="CC" name="Credit-Control">
            <request>
                <!-- 3GPP 29.212 Section 5.6.2 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="Auth-Application-Id" required="true" max="1"/>
                <rule avp="CC-Request-Type" required="true" max="1"/>
                <rule avp="CC-Request-Number" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Origin-State-Id" required="false" max="1"/>
                <rule avp="Subscription-Id" required="false" max="1"/>
                <rule avp="Termination-Cause" required="false" max="1"/>
                <rule avp="User-Equipment-Info" required="false" max="1"/>
                <rule avp="Proxy-Info" required="false" max="1"/>
                <rule avp="Route-Record" required="false" max="1"/>
                <rule avp="Framed-IP-Address" required="false" max="1"/>
                <rule avp="IP-CAN-Type" required="false" max="1"/>
                <rule avp="Called-Station-Id" required="false" max="1"/>
                <rule avp="RAT-Type" required="false" max="1"/>
                <rule avp="Network-Request-Support" required="false" max="1"/>
                <rule avp="Default-EPS-Bearer-QoS" required="false" max="1"/>
                <rule avp="AN-GW-Address" required="false" max="2"/>
                <rule avp="Bearer-Usage" required="false" max="1"/>
                <rule avp="Online" required="false" max="1"/>
                <rule avp="Offline" required="false" max="1"/>
                <rule avp="Access-Network-Charging-Identifier-Gx" required="false"/>
                <rule avp="TGPP-SGSN-Address" required="false" max="1"/>
                <rule avp="TGPP-GGSN-Address" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="Access-Network-Charging-Address" required="false" max="1"/>
                <rule avp="TGPP-MS-TimeZone" required="false" max="1"/>
                <rule avp="TGPP-Selection-Mode" required="false" max="1"/>
                <rule avp="QoS-Information" required="false" max="1"/>
                <rule avp="TGPP-SGSN-MCC-MNC" required="false" max="1"/>
                <rule avp="TGPP-User-Location-Info" required="false" max="1"/>
            </request>
            <answer>
                <!-- 3GPP 29.212 Section 5.6.3 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="CC-Request-Type" required="true" max="1"/>
                <rule avp="CC-Request-Number" required="true" max="1"/>
                <rule avp="Origin-State-Id" required="false" max="1"/>
                <rule avp="Proxy-Info" required="false" max="1"/>
                <rule avp="Route-Record" required="false" max="1"/>
                <rule avp="Failed-AVP" required="false" max="1"/>
                <rule avp="Charging-Rule-Install" required="false"/>
                <rule avp="Charging-Rule-Remove" required="false"/>
                <rule avp="Usage-Monitoring-Information" required="false"/>
                <rule avp="Event-Trigger" required="false"/>
                <rule avp="Revalidation-Time" required="false"/>
            </answer>
        </command>

        <command code="258" short="RA" name="Re-Auth">
            <request>
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="true" max="1"/>
                <rule avp="Auth-Application-Id" required="true" max="1"/>
                <rule avp="Re-Auth-Request-Type" required="true" max="1"/>
                <rule avp="QoS-Information" required="false" max="1"/>
                <rule avp="Origin-State-Id" required="false" max="1"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
                <rule avp="Event-Trigger" required="false"/>
                <rule avp="Revalidation-Time" required="false"/>
            </request>
            <answer>
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Origin-State-Id" required="false" max="1"/>
                <rule avp="Error-Message" required="false" max="1"/>
                <rule avp="Error-Reporting-Host" required="false" max="1"/>
                <rule avp="Failed-AVP" required="false" max="1"/>
                <rule avp="Proxy-Info" required="false"/>
            </answer>
        </command>

        <avp name="Flow-Description" code="507" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="IPFilterRule"/>
        </avp>


        <avp name="Charging-Rule-Install" code="1001" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.2 -->
            <data type="Grouped">
                <rule avp="Charging-Rule-Name" required="false"/>
                <rule avp="Charging-Rule-Base-Name" required="false"/>
                <rule avp="Charging-Rule-Definition" required="false"/>
                <rule avp="Rule-Activation-Time" required="false"/>
                <rule avp="Rule-Deactivation-Time" required="false"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

        <avp name="Charging-Rule-Remove" code="1002" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.3 -->
            <data type="Grouped">
                <rule avp="Charging-Rule-Name" required="false"/>
                <rule avp="Charging-Rule-Base-Name" required="false"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

        <avp name="Charging-Rule-Definition" code="1003" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Grouped">
                <rule avp="Charging-Rule-Name" required="true" max="1"/>
                <rule avp="Rating-Group" required="false" max="1"/>
                <rule avp="Flow-Information" required="false"/>
                <rule avp="Flow-Description" required="false"/>
                <rule avp="Precedence" required="false" max="1"/>
                <rule avp="Monitoring-Key" required="false" max="1"/>
                <rule avp="Redirect-Information" required="false" max="1"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

        <avp name="Charging-Rule-Base-Name" code="1004" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.6 -->
            <data type="UTF8String"/>
        </avp>

        <avp name="Charging-Rule-Name" code="1005" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.6 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Event-Trigger" code="1006" must="M,V" map="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.7 -->
            <data type="Enumerated">
                <item code="0" name="SGSN_CHANGE"/>
                <item code="1" name="QOS_CHANGE"/>
                <item code="2" name="RAT_CHANGE"/>
                <item code="3" name="TFT_CHANGE"/>
                <item code="4" name="PLMN_CHANGE"/>
                <item code="5" name="LOSS_OF_BEARER"/>
                <item code="6" name="RECOVERY_OF_BEARER"/>
                <item code="7" name="IP-CAN_CHANGE"/>
                <item code="11" name="QOS_CHANGE_EXCEEDING_AUTHORIZATION"/>
                <item code="12" name="RAI_CHANGE"/>
                <item code="13" name="USER_LOCATION_CHANGE"/>
                <item code="14" name="NO_EVENT_TRIGGERS"/>
                <item code="15" name="OUT_OF_CREDIT"/>
                <item code="16" name="REALLOCATION_OF_CREDIT"/>
                <item code="17" name="REVALIDATION_TIMEOUT"/>
                <item code="18" name="UE_IP_ADDRESS_ALLOCATE"/>
                <item code="19" name="UE_IP_ADDRESS_RELEASE"/>
                <item code="20" name="DEFAULT_EPS_BEARER_QOS_CHANGE"/>
                <item code="21" name="AN_GW_CHANGE"/>
                <item code="22" name="SUCCESSFUL_RESOURCE_ALLOCATION"/>
                <item code="23" name="RESOURCE_MODIFICATION_REQUEST"/>
                <item code="24" name="PGW_TRACE_CONTROL"/>
                <item code="25" name="UE_TIME_ZONE_CHANGE"/>
                <item code="26" name="TAI_CHANGE"/>
                <item code="27" name="ECGI_CHANGE"/>
                <item code="28" name="CHARGING_CORRELATION_EXCHANGE"/>
                <item code="29" name="APN-AMBR_MODIFICATION_FAILURE"/>
                <item code="30" name="USER_CSG_INFORMATION_CHANGE"/>
                <item code="33" name="USAGE_REPORT"/>
            </data>
        </avp>

        <avp name="Revalidation-Time" code="1042" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212  Section 5.3.41 -->
            <data type="Time"/>
        </avp>

        <avp name="Precedence" code="1010" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Unsigned32"/>
        </avp>

        <avp name="ToS-Traffic-Class" code="1014" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.15 -->
            <data type="Unsigned32"/>
        </avp>

        <avp name="IP-CAN-Type" code="1027" must="M,V" map="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.27 -->
            <data type="Enumerated">
                <item code="0" name="3GPP-GPRS"/>
                <item code="1" name="DOCSIS"/>
                <item code="2" name="xDSL"/>
                <item code="3" name="WiMAX"/>
                <item code="4" name="3GPP2"/>
                <item code="5" name="3GPP-EPS"/>
                <item code="6" name="Non-3GPP-EPS"/>
                <item code="7" name="FBA"/>
                <item code="8" name="3GPP-5GS"/>
                <item code="9" name="Non-3GPP-5GS"/>
            </data>
        </avp>

        <avp name="Rule-Activation-Time" code="1043" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Time"/>
        </avp>

        <avp name="Rule-Deactivation-Time" code="1044" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Time"/>
        </avp>

        <avp name="Security-Parameter-Index" code="1056" must="V" must-not="M" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.51 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Flow-Label" code="1057" must="V" must-not="M" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.52 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Flow-Information" code="1058" must="V" must-not="M" may="P" may-encryp="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.53 -->
            <data type="Grouped">
                <rule avp="Flow-Description" required="false" max="1"/>
                <rule avp="Packet-Filter-Identifier" required="false" max="1"/>
                <rule avp="Packet-Filter-Usage" required="false" max="1"/>
                <rule avp="ToS-Traffic-Class" required="false" max="1"/>
                <rule avp="Security-Parameter-Index" required="false" max="1"/>
                <rule avp="Flow-Label" required="false" max="1"/>
                <rule avp="Flow-Direction" required="false" max="1"/>
            </data>
        </avp>

        <avp name="Packet-Filter-Identifier" code="1060" must="V" must-not="M" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.55 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Monitoring-Key" code="1066" must="V" may="P" must-not="M" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Usage-Monitoring-Information" code="1067" must="V" may="P" must-not="M,V" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Grouped">
                <rule avp="Monitoring-Key" required="false" max="1"/>
                <rule avp="Granted-Service-Unit" required="false" max="2"/>
                <rule avp="Used-Service-Unit" required="false" max="2"/>
                <rule avp="Usage-Monitoring-Level" required="false" max="1"/>
            </data>
        </avp>

        <avp name="Usage-Monitoring-Level" code="1068" must="V" may="P" must-not="M" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 -->
            <data type="Enumerated">
                <item code="0" name="SESSION_LEVEL"/>
                <item code="1" name="PCC_RULE_LEVEL"/>
            </data>
        </avp>

        <avp name="Packet-Filter-Usage" code="1072" must="V" must-not="M" map="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.66 -->
            <data type="Enumerated">
                <item code="1" name="SEND_TO_UE"/>
            </data>
        </avp>

        <avp name="Flow-Direction" code="1080" must="V" must-not="M" map="P" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.65 -->
            <data type="Enumerated">
                <item code="0" name="UNSPECIFIED"/>
                <item code="1" name="DOWNLINK"/>
                <item code="2" name="UPLINK"/>
                <item code="3" name="BIDIRECTIONAL"/>
            </data>
        </avp>

        <avp name="Redirect-Information" code="1085" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.82 -->
            <data type="Grouped">
                <rule avp="Redirect-Support" required="true" max="1"/>
                <rule avp="Redirect-Address-Type" required="false" max="1"/>
                <rule avp="Redirect-Server-Address" required="false" max="1"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

        <avp name="Redirect-Support" code="1086" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.83 -->
            <data type="Enumerated">
                <item code="0" name="REDIRECTION_DISABLED"/>
                <item code="1" name="REDIRECTION_ENABLED"/>
            </data>
        </avp>

        <avp name="Network-Request-Support" code="1024" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.24 -->
            <data type="Enumerated">
                <item code="0" name="NETWORK_REQUEST_NOT_SUPPORTED"/>
                <item code="1" name="NETWORK_REQUEST_SUPPORTED"/>
            </data>
        </avp>

        <avp name="Offline" code="1008" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.9 -->
            <data type="Enumerated">
                <item code="0" name="DISABLE_OFFLINE"/>
                <item code="1" name="ENABLE_OFFLINE"/>
            </data>
        </avp>

        <avp name="Online" code="1009" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.10 -->
            <data type="Enumerated">
                <item code="0" name="DISABLE_ONLINE"/>
                <item code="1" name="ENABLE_ONLINE"/>
            </data>
        </avp>

        <avp name="Default-EPS-Bearer-QoS" code="1049" must="V" may="P" must-not="M" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.48 -->
            <data type="Grouped">
                <rule avp="QoS-Class-Identifier" required="false" max="1"/>
                <rule avp="Allocation-Retention-Priority" required="false" max="1"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

        <avp name="AN-GW-Address" code="1050" must="V" may="P" must-not="M" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.49 -->
            <data type="Address"/>
        </avp>

        <avp name="Bearer-Usage" code="1000" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.1 -->
            <data type="Enumerated">
                <item code="0" name="GENERAL"/>
                <item code="1" name="IMS_SIGNALLING"/>
            </data>
        </avp>

        <avp name="Access-Network-Charging-Identifier-Gx" code="1022" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 Section 5.3.22 -->
            <data type="Grouped">
                <rule avp="Access-Network-Charging-Identifier-Value" required="true" max="1"/>
                <rule avp="Charging-Rule-Base-Name" required="false"/>
                <rule avp="Charging-Rule-Name" required="false"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

        <avp name="TGPP-SGSN-Address" code="6" must="V" may="P" must-not="M" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.061 Table 9a -->
            <data type="OctetString"/>
        </avp>

        <avp name="TGPP-GGSN-Address" code="7" must="V" may="P" must-not="M" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.061 Table 9a -->
            <data type="OctetString"/>
        </avp>

        <avp name="Access-Network-Charging-Address" code="501" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP TS 29.214 Section 5.3.2 -->
            <data type="Address"/>
        </avp>

        <avp name="TGPP-MS-TimeZone" code="23" must="V" may="P" must-not="M" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.061 Table 9a -->
            <data type="OctetString"/>
        </avp>

        <avp name="TFT-Filter" code="1012" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 5.3.13-->
            <data type="IPFilterRule"/>
        </avp>

        <avp name="TFT-Packet-Filter-Information" code="1013" must="M,V" may="P" may-encrypt="y" vendor-id="10415">
            <!-- 3GPP 29.212 5.3.14-->
            <data type="Grouped">
                <rule avp="Precedence" required="false" max="1"/>
                <rule avp="TFT-Filter" required="false" max="1"/>
                <rule avp="ToS-Traffic-Class" required="false" max="1"/>
                <rule avp="Security-Parameter-Index" required="false" max="1"/>
                <rule avp="Flow-Label" required="false" max="1"/>
                <rule avp="Flow-Direction" required="false" max="1"/>
                <!-- *[ AVP ]-->
            </data>
        </avp>

    </application>
</diameter>`

var tgppe2XML = `<?xml version="1.0" encoding="UTF-8"?>
<diameter>
        <!--
        3GPP TS 283 035
        See: https://www.etsi.org/deliver/etsi_es/283000_283099/283035/03.02.01_60/es_283035v030201p.pdf
    -->
    <application id="16777231" type="auth" name="TGPP e2">
         <vendor id="10415" name="TGPP"/>
         <command code="306" short="UD" name="User-Data">
            <!--
                < User-Data -Request > ::= < Diameter Header: 306, REQ, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [ Destination-Host ]
                { Destination-Realm }
                [Globally-Unique-Address]
                [User-Name]
                [Global-Access-Id]
                [AF-Application-Identifier]
                [Requested-Information]
                *[ AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <request>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="Destination-Host" required="false" max="1" />
                <rule avp="Destination-Realm" required="true" max="1" />
                <rule avp="Globally-Unique-Address" required="false" max="1" />
                <rule avp="User-Name" required="false" max="1" />
                <rule avp="Global-Access-Id" required="false" max="1" />
                <rule avp="AF-Application-Identifier" required="false" max="1" />
                <rule avp="Requested-Information" required="false" max="1" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />
            </request>

            <!--
                < User-Data-Answer > ::= < Diameter Header: 306, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                [ Result-Code ]
                [ Experimental-Result ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [User-Name]
                [Logical-Access-Id]
                [Physical-Access-Id]
                [Access-Network-Type]
                [Location-Information]
                [Policy-Control-Contact-Point]
                [Terminal-Type]
                *[ AVP ]
                *[ Failed-AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <answer>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1" />
                <rule avp="Result-Code" required="false" max="1" />
                <rule avp="Experimental-Result" required="false" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="User-Name" required="false" max="1" />
                <rule avp="Logical-Access-Id" required="false" max="1" />
                <rule avp="Physical-Access-Id" required="false" max="1" />
                <rule avp="Access-Network-Type" required="false" max="1" />
                <rule avp="Location-Information" required="false" max="1" />
                <rule avp="Policy-Control-Contact-Point" required="false" max="1" />
                <rule avp="Terminal-Type" required="false" max="1" />
                <rule avp="AVP" required="false"/>
                <rule avp="Failed-AVP" required="false" max="1" />
                <rule avp="Proxy-Info" required="false" />
                <rule avp="Route-Record" required="false" />
            </answer>
         </command>

         <command code="308" short="SN" name="Subscribe-Notifications">

            <!--
                < Subscribe-Notifications-Request > ::= < Diameter Header: 308, REQ, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                [ Destination-Host ]
                { Destination-Realm }
                { Subs-Req-Type }
                [ Expiry-Time ]
                [Globally-Unique-Address]
                [User-Name]
                [AF-Application-Identifier]
                *[Event-Type]
                *[ AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <request>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="Destination-Host" required="false" max="1" />
                <rule avp="Destination-Realm" required="true" max="1" />
                <rule avp="Subs-Req-Type" required="true" max="1" />
                <rule avp="Expiry-Time" required="false" max="1" />
                <rule avp="Globally-Unique-Address" required="false" max="1" />
                <rule avp="User-Name" required="false" max="1"/>
                <rule avp="AF-Application-Identifier" required="false" max="1" />
                <rule avp="Event-Type" required="false" max="1" />
                <rule avp="AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </request>

            <!--
                < Subscribe-Notifications-Answer > ::= < Diameter Header: 308, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                { Auth-Session-State }
                [ Result-Code ]
                [ Experimental-Result ]
                { Origin-Host }
                { Origin-Realm }
                [ Expiry-Time ]
                *[ AVP ]
                *[ Failed-AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <answer>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Result-Code" required="false" max="1" />
                <rule avp="Experimental-Result" required="false" max="1" />
                <rule avp="Origin-Host" required="false" max="1" />
                <rule avp="Origin-Realm" required="false" max="1" />
                <rule avp="Expiry-Time" required="false" max="1" />
                <rule avp="AVP" required="false"/>
                <rule avp="Failed-AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </answer>
         </command>

         <command code="309" short="PN" name="Push-Notification">

            <!--
                < Push-Notification-Request > ::= < Diameter Header: 309, REQ, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                { Destination-Host }
                { Destination-Realm }
                *[Event-Type]
                [Globally-Unique-Address]
                [User-Name]
                [Location-Information]
                [Policy-Control-Contact-Point]
                [Terminal-Type]
                [Logical-Access-Id]
                [Physical-Access-Id]
                [Access-Network-Type]
                [Initial-Gate-Setting]
                *[QoS-Profile]
                [IP-Connectivity-Status]
                *[ AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <request>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Origin-Host" required="true" max="1" />
                <rule avp="Origin-Realm" required="true" max="1" />
                <rule avp="Destination-Host" required="true" max="1" />
                <rule avp="Destination-Realm" required="true" max="1" />
                <rule avp="Event-Type" required="false" max="1" />
                <rule avp="Globally-Unique-Address" required="false" max="1" />
                <rule avp="User-Name" required="false" max="1" />
                <rule avp="Location-Information" required="false" max="1" />
                <rule avp="Policy-Control-Contact-Point" required="false" max="1" />
                <rule avp="Terminal-Type" required="false" max="1" />
                <rule avp="Logical-Access-ID" required="false" max="1" />
                <rule avp="Physical-Access-ID" required="false" max="1" />
                <rule avp="Access-Network-Type" required="false" max="1" />
                <rule avp="Initial-Gate-Setting" required="false" max="1" />
                <rule avp="QoS-Profile" required="false" max="1" />
                <rule avp="IP-Connectivity-Status" required="false" max="1" />
                <rule avp="AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </request>

            <!--
                < Push-Notification-Answer > ::= < Diameter Header: 309, PXY, 16777231 >
                < Session-Id >
                { Vendor-Specific-Application-Id }
                [ Result-Code ]
                [ Experimental-Result ]
                { Auth-Session-State }
                { Origin-Host }
                { Origin-Realm }
                *[ AVP ]
                *[ Failed-AVP ]
                *[ Proxy-Info ]
                *[ Route-Record ]
            -->
            <answer>
                <rule avp="Session-Id" required="true" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1" />
                <rule avp="Auth-Session-State" required="true" max="1" />
                <rule avp="Result-Code" required="false" max="1" />
                <rule avp="Experimental-Result" required="false" max="1" />
                <rule avp="Origin-Host" required="false" max="1" />
                <rule avp="Origin-Realm" required="false" max="1" />
                <rule avp="AVP" required="false"/>
                <rule avp="Failed-AVP" required="false"/>
                <rule avp="Proxy-Info" required="false"/>
                <rule avp="Route-Record" required="false"/>
            </answer>
         </command>

        <avp name="Location-Information" code="350" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Line-Identifier" required="false" max="1" />
                <rule avp="Civic-Location" required="false" max="1" />
                <rule avp="Geospatial-Location" required="false" max="1" />
                <rule avp="Operator-Specific-GI" required="false" max="1" />
                <rule avp="PIDF-Location-Object" required="false" max="1" />
                <rule avp="AVP" required="false" />
            </data>
        </avp>

        <avp name="Policy-Control-Contact-Point" code="351" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="DiameterIdentity"/>
        </avp>

        <avp name="Terminal-Type" code="352" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Requested-Information" code="353" vendor-id="13019" must="V" must-not="M" may-encrypt="Y">
            <data type="Enumerated">
                <item code="0" name="NASS-USER-ID"/>
                <item code="1" name="LOCATION-INFORMATION"/>
                <item code="2" name="POLICY-CONTROL-CONTACT-POINT"/>
                <item code="3" name="ACCESS-NETWORK-TYPE"/>
                <item code="4" name="TERMINAL-TYPE"/>
                <item code="5" name="LOGICAL-ACCESS-ID"/>
                <item code="6" name="PHYSICAL-ACCESS-ID"/>
                <item code="7" name="IP-ADDRESS"/>
                <item code="8" name="INITIAL-GATE-SETTING"/>
                <item code="9" name="QOS-PROFILE"/>
                <item code="10" name="IP-CONNECTIVITY-STATUS"/>
                <item code="11" name="EMERGENCY-CALL-ROUTING-INFO"/>
            </data>
        </avp>

        <avp name="Event-Type" code="354" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="Enumerated">
                <item code="0" name="USER-LOGON"/>
                <item code="1" name="LOCATION-INFORMATION-CHANGED"/>
                <item code="2" name="POLICY-CONTROL-CONTACT-POINT-CHANGED"/>
                <item code="3" name="ACCESS-NETWORK-TYPE-CHANGED"/>
                <item code="4" name="TERMINAL-TYPE-CHANGED"/>
                <item code="5" name="LOGICAL-ACCESS-ID-CHANGED"/>
                <item code="6" name="PHYSICAL-ACCESS-ID-CHANGED"/>
                <item code="7" name="IP-ADDRESS-CHANGED"/>
                <item code="8" name="INITIAL-GATE-SETTING-CHANGED"/>
                <item code="9" name="QOS-PROFILE-CHANGED"/>
                <item code="10" name="USER-LOGOFF"/>
            </data>
        </avp>

        <avp name="Line-Identifier" code="500" vendor-id="13019" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Civic-Location" code="355" vendor-id="13019" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Geospatial-Location" code="356" vendor-id="13019" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Global-Access-Id" code="357" vendor-id="13019" must="V" must-not="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Fixed-Access-ID" required="false" max="1" />
                <rule avp="3GPP-User-Location-Info" required="false" max="1" />
                <rule avp="AVP" required="false" />
            </data>
        </avp>

        <avp name="Fixed-Access-Id" code="358" vendor-id="13019" must="V" must-not="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Relay-Agent" required="true" max="1" />
                <rule avp="Logical-Access-ID" required="false" max="1" />
                <rule avp="Physical-Access-ID" required="false" max="1" />
                <rule avp="AVP" required="false" />
            </data>
        </avp>

        <avp name="Relay-Agent" code="359" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Operator-Specific-GI" code="360" vendor-id="13019" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Emergency-Call-Routing-Info" code="361" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="UTF8String"/>
        </avp>

        <avp name="Port-Number" code="362" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="Unsigned32"/>
        </avp>

        <avp name="PIDF-Location-Object" code="363" vendor-id="13019" must="V" must-not="M" may-encrypt="Y">
            <data type="UTF8String"/>
        </avp>

        <avp name="AF-Application-Identifier" code="504" vendor-id="10415" must="M,V" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Globally-Unique-Address" code="300" vendor-id="13019" must="M,V" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Framed-IP-Address" required="false" max="1" />
                <rule avp="Framed-IPv6-Prefix" required="false" max="1" />
                <rule avp="Address-Realm" required="false" max="1" />
            </data>
        </avp>

        <avp name="Logical-Access-Id" code="302" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Access-Network-Type" code="306" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="NAS-Port-Type" required="true" max="1" />
                <rule avp="Aggregation-Network-Type" required="false" max="1" />
            </data>
        </avp>

        <avp name="Initial-Gate-Setting" code="303" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="NAS-Filter-Rule" required="false" max="1" />
                <rule avp="Max-Requested-Bandwidth-UL" required="false" max="1" />
                <rule avp="Max-Requested-Bandwidth-DL" required="false" max="1" />
            </data>
        </avp>

        <avp name="QoS-Profile" code="304" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="Grouped">
                <rule avp="Application-Class-ID" required="false" max="1" />
                <rule avp="Media-Type" required="false" max="1" />
                <rule avp="Reservation-Priority" required="false" max="1" />
                <rule avp="Max-Requested-Bandwidth-UL" required="false" max="1" />
                <rule avp="Max-Requested-Bandwidth-DL" required="false" max="1" />
                <rule avp="Transport-Class" required="false" max="1" />
            </data>
        </avp>

        <avp name="IP-Connectivity-Status" code="305" vendor-id="13019" must="V" may="M" may-encrypt="Y">
            <data type="Enumerated">
                <item code="0" name="IP-CONNECTIVITY-ON"/>
                <item code="1" name="IP-CONNECTIVITY-LOST"/>
            </data>
        </avp>

        <avp name="Aggregation-Network-Type" code="307" vendor-id="10415" must="V" may="M" may-encrypt="Y">
            <data type="Enumerated">
                <item code="0" name="UNKNOWN"/>
                <item code="1" name="ATM"/>
                <item code="2" name="ETHERNET"/>
            </data>
        </avp>

        <avp name="Expiry-Time" code="709" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="Time">
            </data>
        </avp>

        <avp name="Subs-Req-Type" code="705" vendor-id="10415" must="M,V" may-encrypt="Y">
            <data type="Enumerated">
                <item code="0" name="Subscribe"/>
                <item code="1" name="Unsubscribe"/>
            </data>
        </avp>

        <avp name="TGPP-User-Location-Info" code="22" vendor-id="10415" must="V" must-not="M" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Framed-IP-Address" code="8" may="M" must-not="V" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Framed-IPv6-Prefix" code="97" may="M" must-not="V" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Address-Realm" code="301" vendor-id="13015" may="M,V" may-encrypt="Y">
            <data type="OctetString"/>
        </avp>

        <avp name="Max-Requested-Bandwidth-DL" code="515" must="V,M"    may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>

        <avp name="Max-Requested-Bandwidth-UL" code="516" must="V,M"    may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
            <data type="Unsigned32"/>
        </avp>


    </application>
</diameter>`
