//SPDX-License-Identifier: TEST_EXCHANGE
pragma solidity ^0.8.4;

// Logger contract used to anonymously log state transitions of orders and trades
contract Logger {

    // Possible states of an order
    enum TStateOrder {
        CREATED,
        EXPIRED,
        TRADED
    }

    // Possible states of a trade
    enum TStateTrade {
        CREATED,
        SIGNATURES,
        INSTRUCTIONS,
        SHIPPINGADVICE,
        DOCUMENTS,
        COMPLETION,
        INTENTION,
        DECLARATION,
        CLOSED
    }

    // Type of transaction (ORDER|OTCORDER|TRADE)
    enum TType {
        ORDER,
        OTCORDER,
        TRADE
    }

    // Data structure of the logged transaction detail
    struct TLog {
        string hash;
        TType typeInd;
        TStateOrder stOrder;
        TStateTrade stTrade;
        uint256 timestamp;
    }

    // Data structure for historical state record
    struct TStateLog {
        uint8 tt;
        uint8 st;
        uint256 timestamp;
    }

    // Function to retrieve all historical states with timestamps of an order/trade
    function getHistoricStates(bytes32 _identifier, TType _tp) view public returns (TStateLog[] memory){

        // Validates Type parameter that was used
        require(
            _tp == TType.ORDER || _tp == TType.OTCORDER || _tp == TType.TRADE,
            "Type is not correct."
        );

        // Return the history collected for the type specified
        if ( _tp == TType.ORDER ) {
            return oStatesList[_identifier];
        } else if ( _tp == TType.OTCORDER ) {
            return otcStatesList[_identifier];
        } else {
            return tStatesList[_identifier];
        }
    }

    // Function to retrieve just the hash value of a logged transaction identified by _identifier
    function getHash(bytes32 _identifier, TType _tp) view public returns (string memory){

        // Validates Type parameter that was used
        require(
            _tp == TType.ORDER || _tp == TType.OTCORDER || _tp == TType.TRADE,
            "Type is not correct."
        );

        // Return the history collected for the type specified
        if ( _tp == TType.ORDER ) {
            return( oLogList[_identifier].hash );
        } else if ( _tp == TType.OTCORDER ) {
            return( otcLogList[_identifier].hash );
        } else {
            return( tLogList[_identifier].hash );
        }
    }

    // Function to retrieve all log details of a logged transaction identified by _identifier
    function getLog(bytes32 _identifier, TType _tp) view public returns (string memory, TType, uint8, uint256)  {


        // Validates Type parameter that was used
        require(
            _tp == TType.ORDER || _tp == TType.OTCORDER || _tp == TType.TRADE,
            "Type is not correct."
        );

        // depending on the Type return a values holding the right state
        if (_tp == TType.ORDER)  {
            TLog memory log = oLogList[_identifier];
            return (
            log.hash,
            log.typeInd,
            uint8(log.stOrder),
            log.timestamp
            );
        } else if (_tp == TType.OTCORDER) {
            TLog memory log = otcLogList[_identifier];
            return (
            log.hash,
            log.typeInd,
            uint8(log.stOrder),
            log.timestamp
            );
        } else {
            TLog memory log = tLogList[_identifier];
            return(
            log.hash,
            log.typeInd,
            uint8(log.stTrade),
            log.timestamp
            );
        }
    }

    // New LogLists one for each type (this to avoid key-collisions between the types)
    mapping(bytes32 => TLog) private oLogList;
    mapping(bytes32 => TLog) private otcLogList;
    mapping(bytes32 => TLog) private tLogList;

    // New StatesLists one for each type (this to avoid key-collisions between the types)
    mapping(bytes32 => TStateLog[]) private oStatesList;
    mapping(bytes32 => TStateLog[]) private otcStatesList;
    mapping(bytes32 => TStateLog[]) private tStatesList;

    // var to hold the owner address
    address owner;

    // constructor sets owner of transaction
    constructor() {
        owner = msg.sender;
    }

    // Log transaction hash and details; ID, blocktime, State and Type
    function hashLog(
        bytes32 id,
        string memory hash,
        TType tp,
        uint8 st
    ) public {
        // caller must be the owner of this contract
        require(msg.sender == owner);

        // Validates Type parameter that was used
        require(
            tp == TType.ORDER || tp == TType.OTCORDER || tp == TType.TRADE,
            "Type is not correct."
        );

        // New Check to confirm that state is higher or same tha previous state ORDER-OTCORDER-TRADE
        if (tp == TType.ORDER && oStatesList[id].length > 0) {

            // ORDER: New Check to confirm that order state is higher or same than previous state
            TStateLog memory oPrev = oStatesList[id][oStatesList[id].length-1];
            require(oPrev.st <= st, "The new order state must be higher or same then previous.");
        } else if (tp == TType.OTCORDER && otcStatesList[id].length > 0) {

            // OTCORDER: New Check to confirm that OTC-order state is higher or same than previous state
            TStateLog memory otcPrev = otcStatesList[id][otcStatesList[id].length-1];
            require(otcPrev.st <= st, "The new OTC-order state must be higher or same then previous.");

        } else if (tp == TType.TRADE && tStatesList[id].length > 0) {

            // TRADE: New Check to confirm that trade state is higher or same than previous state
            TStateLog memory tPrev = tStatesList[id][tStatesList[id].length-1];
            require(tPrev.st <= st, "The new trade state must be higher or same then previous.");
        }

        // Get block-time and save as BigInt
        uint256 time = block.timestamp;

        // Tests to see if state is legal for the type of transaction that we are logging
        if (tp == TType.TRADE) {

            require(st <= 8, "Trade state is not correct.");
            TLog memory log = TLog(hash, tp, TStateOrder(0), TStateTrade(st), time);
            tLogList[id] = log;

            // Add historical state + time tuple to historical States for this trade with id
            tStatesList[id].push(TStateLog(uint8(tp), st, time));

        } else if (tp == TType.OTCORDER) {
            require(st <= 2, "OTC-order state is not correct.");
            TLog memory log = TLog(hash, tp, TStateOrder(st), TStateTrade(0), time);
            otcLogList[id] = log;

            // Add historical state + time tuple to historical States for this trade with id
            otcStatesList[id].push(TStateLog(uint8(tp), st, time));
        } else {
            require(st <= 2, "Order state is not correct.");
            TLog memory log = TLog(hash, tp, TStateOrder(st), TStateTrade(0), time);
            oLogList[id] = log;

            // Add historical state + time tuple to historical States for this trade with id
            oStatesList[id].push(TStateLog(uint8(tp), st, time));
        }
    }
}
